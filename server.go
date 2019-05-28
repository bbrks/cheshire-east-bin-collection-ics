package main

import (
	"errors"
	"fmt"
	"io"
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
}

func (s *server) handleCollections() http.HandlerFunc {
	collectionDates := func(r io.Reader) (*Collections, error) {
		collections := NewCollections()

		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			return nil, err
		}

		jobList := doc.Find("table.full-job-list-for-posting")
		if jobList.Length() == 0 {
			return nil, errors.New("No collection list found in HTML")
		} else if jobList.Length() > 1 {
			s.Log(LevelWarn, nil, "More than one collection list found in HTML... using first one only.")
		}

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
				s.Log(LevelWarn, nil, "Some data was unexpectedly missing in form")
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
				s.Log(LevelWarn, nil, "Unexpected class for given bin %v %v", bin, class)
			}

			// Format: 14/05/2019 07:00:00
			date, err := time.Parse("02/01/2006 15:04:05", cScheduledStart)
			if err != nil {
				s.Log(LevelError, nil, "Invalid date: %v", cScheduledStart)
			}

			collections.Add(Collection{
				ID:   cId,
				Bin:  bin,
				Time: date,
				Name: cName,
				Description: cJobDescription,
			})
		})

		return collections, nil
	}

	// fetch gets the current collection dates for the given UPRN
	fetch := func(uprn string) (*Collections, error) {
		resp, err := http.Get(fetchURL + uprn)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("error fetching data: %v", err)
		}

		return collectionDates(resp.Body)
	}

	//fetchStub := func(uprn string) (*Collections, error) {
	//	f, err := os.Open("fixtures/source-data.html")
	//	if err != nil {
	//		return nil, err
	//	}
	//	defer f.Close()
	//
	//	return collectionDates(f)
	//}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "text/calendar")
		w.Header().Set("charset", "utf-8")
		w.Header().Set("Content-Disposition", "inline")
		w.Header().Set("filename", calendarName)

		collections, err := fetch(s.uprn)
		if err != nil {
			write(w, http.StatusInternalServerError, []byte(err.Error()))
			return
		}

		goics.NewICalEncode(w).Encode(collections)
	}
}

func write(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	_, _ = w.Write(body)
	_, _ = w.Write([]byte("\n"))
}
