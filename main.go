package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	appName    = "cheshire-east-bin-collection-ics"
	appVersion = "0.1"

	fetchURL     = "https://online.cheshireeast.gov.uk/MyCollectionDay/SearchByAjax/GetBartecJobList?uprn="
	calendarName = "collections.ics"
)

func main() {
	envUPRN := os.Getenv("UPRN")

	flagUPRN := flag.String("uprn", "", "Your Unique Property Reference Number (UPRN)")
	flagUpdateInterval := flag.Duration("updateInterval", time.Hour*24, "Data younger than this value will be served from the cache")
	flagBindAddr := flag.String("http", ":8080", "The address and/or port to bind to the HTTP server")
	flagLogQuiet := flag.Bool("q", false, "Disables all logging except for errors")
	flag.Parse()

	s := server{
		Logger:         newDefaultLogger(*flagLogQuiet),
		router:         httprouter.New(),
		httpClient:     http.DefaultClient,
		updateInterval: *flagUpdateInterval,
		cache:          NewCollections(),
		cachedAt:       time.Time{},
	}
	ctx := context.Background()

	uprn := *flagUPRN
	if uprn == "" {
		uprn = envUPRN
	}
	if uprn == "" {
		s.Log(LevelError, ctx, "UPRN was not set. Set with \"UPRN\" env var or -uprn=\"your-uprn\" flag")
		os.Exit(1)
	}

	if *flagUpdateInterval < time.Hour*24 {
		s.Log(LevelError, ctx, "Invalid updateInterval set. Must be at least 24h")
		os.Exit(1)
	} else if *flagUpdateInterval > time.Hour*24*7 {
		s.Log(LevelError, ctx, "Invalid updateInterval set. Must be at most 168h (7 days)")
		os.Exit(1)
	}

	s.uprn = uprn
	s.routes()

	l, err := net.Listen("tcp", *flagBindAddr)
	if err != nil {
		s.Log(LevelError, ctx, err.Error())
		os.Exit(1)
	}

	s.Log(LevelAll, ctx, "Serving calendar at http://%s/%s", l.Addr().String(), calendarName)
	err = http.Serve(l, s.middleware(s.router.ServeHTTP))
	if err != nil {
		s.Log(LevelError, ctx, "Error from HTTP server: %v", err)
		os.Exit(1)
	}
}
