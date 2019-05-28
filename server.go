package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jordic/goics"
	"github.com/julienschmidt/httprouter"
)

type server struct {
	Logger     // Embed Log()
	router     *httprouter.Router
	httpClient *http.Client

	uprn           string
	updateInterval time.Duration
	cache          *Collections
	cachedAt       time.Time
}

var errNoData = errors.New("no collection data found in HTML")

func (s *server) handleCollections(fetchFn func(uprn string) (io.ReadCloser, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if time.Since(s.cachedAt) > s.updateInterval {
			s.Log(LevelInfo, r.Context(), "Fetching a fresh copy of data")
			readCloser, err := fetchFn(s.uprn)
			if err != nil {
				s.Log(LevelError, r.Context(), "Error fetching data: %v - serving stale cache...", err)
				serveCollections(w, s.cache)
				return
			}
			defer readCloser.Close()

			parsedCollections, err := s.parseCollections(r.Context(), readCloser)
			if err == errNoData {
				s.Log(LevelError, r.Context(), "No data found for UPRN: %v", err)
				write(w, http.StatusBadRequest, []byte(err.Error()))
				return
			} else if err != nil {
				s.Log(LevelError, r.Context(), "Error parsing data: %v - serving stale cache...", err)
				serveCollections(w, s.cache)
				return
			}

			s.cache = parsedCollections
			s.cachedAt = time.Now()
		} else {
			s.Log(LevelDebug, r.Context(), "Skipping update and serving data from cache")
		}

		serveCollections(w, s.cache)
	}
}

func serveCollections(w http.ResponseWriter, c *Collections) {
	if c.Len() == 0 {
		write(w, http.StatusBadRequest, []byte("No collection data found"))
		return
	}

	w.Header().Set("Content-type", "text/calendar")
	w.Header().Set("charset", "utf-8")
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("filename", calendarName)
	w.WriteHeader(http.StatusOK)
	goics.NewICalEncode(w).Encode(c)
}

func (s *server) parseCollections(ctx context.Context, r io.Reader) (*Collections, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	jobList := doc.Find("table.full-job-list-for-posting")
	if jobList.Length() == 0 {
		return nil, errNoData
	} else if jobList.Length() > 1 {
		s.Log(LevelWarn, ctx, "More than one collection list found in HTML... using first one only.")
	}

	c := NewCollections()

	jobList.First().Find("tr.data-row").Each(func(collectionNum int, sel *goquery.Selection) {
		classBlack := sel.HasClass("non-recyclable-bin-type")
		classSilver := sel.HasClass("recyclable-bin-type")
		classGreen := sel.HasClass("green-waste-bin-type")

		cId, cIdOK := sel.Find(`input[name="BartecSimplifiedJobList[` + strconv.Itoa(collectionNum) + `].Id"]`).First().Attr("value")
		cName, cNameOK := sel.Find(`input[name="BartecSimplifiedJobList[` + strconv.Itoa(collectionNum) + `].Name"]`).First().Attr("value")
		cScheduledStart, cScheduledStartOK := sel.Find(`input[name="BartecSimplifiedJobList[` + strconv.Itoa(collectionNum) + `].ScheduledStart"]`).First().Attr("value")
		cJobDescription, cJobDescriptionOK := sel.Find(`input[name="BartecSimplifiedJobList[` + strconv.Itoa(collectionNum) + `].JobDescription"]`).First().Attr("value")
		_, cRescheduledOK := sel.Find(`input[name="BartecSimplifiedJobList[` + strconv.Itoa(collectionNum) + `].IsRescheduled"]`).First().Attr("value")

		if !cIdOK || !cNameOK || !cScheduledStartOK || !cJobDescriptionOK || !cRescheduledOK {
			s.Log(LevelWarn, ctx, "Some data was unexpectedly missing in form")
		}

		var bin Bin
		var classMismatch bool
		switch cName {
		case "Empty Standard General Waste":
			bin = GeneralWaste
			if !classBlack || classSilver || classGreen {
				classMismatch = true
			}
		case "Empty Standard Mixed Recycling":
			bin = MixedRecycling
			if classBlack || !classSilver || classGreen {
				classMismatch = true
			}
		case "Empty Standard Garden Waste":
			bin = GardenWaste
			if classBlack || classSilver || !classGreen {
				classMismatch = true
			}
		}

		if classMismatch {
			class, _ := sel.Attr("class")
			s.Log(LevelWarn, ctx, "Unexpected class for given bin %v %v", bin, class)
		}

		// Format: 14/05/2019 07:00:00
		date, err := time.Parse("02/01/2006 15:04:05", cScheduledStart)
		if err != nil {
			s.Log(LevelError, ctx, "Invalid date: %v", cScheduledStart)
		}

		c.Add(Collection{
			ID:          cId,
			Bin:         bin,
			Time:        date,
			Name:        cName,
			Description: cJobDescription,
		})
	})

	return c, nil
}

// fetchSample returns a sample collection
func fetchSample(_ string) (io.ReadCloser, error) {
	buf := bytes.NewBufferString(sampleData)
	return ioutil.NopCloser(buf), nil
}

// fetchRemote returns collection dates for the given UPRN from the upstream data source
func fetchRemote(uprn string) (io.ReadCloser, error) {
	resp, err := http.Get(fetchURL + uprn)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error fetching data: %s", string(b))
	}

	return resp.Body, nil
}

func write(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	_, _ = w.Write(body)
	_, _ = w.Write([]byte("\n"))
}
