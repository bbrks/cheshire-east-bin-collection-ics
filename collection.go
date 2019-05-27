package main

import (
	"sort"
	"sync"
	"time"

	"github.com/jordic/goics"
)

type Collection struct {
	ID   string
	Time time.Time
	Bin  Bin
	Name string
	Description string
}

type Collections struct {
	// List of collections
	list []Collection
	// Map of ID to list element lookup
	idIdx map[string]int
	sync.Mutex
}

func (c *Collections) EmitICal() goics.Componenter {
	/*
		BEGIN:VCALENDAR
		VERSION:2.0
		PRODID:-//hacksw/handcal//NONSGML v1.0//EN
			...
		END:VCALENDAR
	*/
	comp := goics.NewComponent()
	comp.SetType("VCALENDAR")
	comp.AddProperty("VERSION", "2.0")
	comp.AddProperty("PRODID", "-//bbrks.me/"+appName+" "+appVersion+"//EN")

	sort.Sort(c)
	for _, collection := range c.List() {
		s := goics.NewComponent()
		s.SetType("VEVENT")

		s.AddProperty("UID", collection.ID)
		s.AddProperty("DESCRIPTION", collection.Name+" "+collection.Description)
		s.AddProperty("SUMMARY", collection.Bin.String())
		s.AddProperty("COLOR", collection.Bin.Color())

		k, v := goics.FormatDateTimeField("DTSTAMP", collection.Time)
		s.AddProperty(k, v)

		k, v = goics.FormatDateTimeField("DTSTART", collection.Time)
		s.AddProperty(k, v)

		k, v = goics.FormatDateTimeField("DTEND", collection.Time.Add(time.Hour * 2))
		s.AddProperty(k, v)

		comp.AddComponent(s)
	}

	return comp
}

var _ goics.ICalEmiter = &Collections{}

var _ sort.Interface = &Collections{}

func NewCollections() *Collections {
	return &Collections{
		list: make([]Collection, 0),
		idIdx: make(map[string]int, 0),
	}
}

func (c *Collections) List() []Collection {
	return c.list
}

func (c *Collections) Add(collection Collection) {
	c.Lock()
	defer c.Unlock()
	newIdx := len(c.list)
	c.list = append(c.list, collection)
	c.idIdx[collection.ID] = newIdx
}

func (c *Collections) Len() int      { return len(c.list) }
func (c *Collections) Swap(i, j int) { c.list[i], c.list[j] = c.list[j], c.list[i] }
func (c *Collections) Less(i, j int) bool {
	if c.list[i].Time.Equal(c.list[j].Time) {
		// If dates match, sort based on Bin type
		return c.list[i].Bin < c.list[j].Bin
	}
	return c.list[i].Time.Before(c.list[j].Time)
}

type Date struct {
	year  int
	month time.Month
	day   int
}

// Compare returns:
// -1 if a < b
//  0 if a == b
// +1 if a > b
func (a Date) Compare(b Date) int {
	// a < b
	if a.year < b.year ||
		a.month < b.month ||
		a.day < b.day {
		return -1
	}

	// a == b
	if a.year == b.year &&
		a.month == b.month &&
		a.day == b.day {
		return 0
	}

	// a > b
	return 1
}
