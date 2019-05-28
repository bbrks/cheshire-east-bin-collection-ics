package main

import (
	"net/http"
)

func (s *server) routes() {
	s.GET("/"+calendarName, s.handleCollections(fetchRemote))
}

func (s *server) GET(path string, h http.HandlerFunc) {
	s.router.HandlerFunc(http.MethodGet, path, h)
}
