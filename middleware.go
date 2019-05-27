package main

import (
	"context"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

// middleware is a set of common middleware handlers for all requests.
func (s *server) middleware(next http.HandlerFunc) http.HandlerFunc {
	return s.mwReqID(s.mwRecover(s.mwServerHeader(s.mwLog(next))))
}

// mwServerHeader adds a Server header to responses.
func (s *server) mwServerHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", appName+"/"+appVersion)
		next.ServeHTTP(w, r)
	}
}

// mwRecover recovers from a panic that occurred when handling a request.
func (s *server) mwRecover(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				s.Log(LevelError, r.Context(), "panic serving request: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
				return
			}
		}()
		next.ServeHTTP(w, r)
	}
}

type reqIDKey struct{}

// mwReqID assigns a unique ID to each HTTP request.
func (s *server) mwReqID(next http.HandlerFunc) http.HandlerFunc {
	// gReqID is an atomically incremented variable to identify each HTTP request with a unique ID.
	var gReqID uint64

	return func(w http.ResponseWriter, r *http.Request) {
		// increment request sequence
		reqID := strconv.FormatUint(atomic.AddUint64(&gReqID, 1), 10)
		r = r.WithContext(context.WithValue(r.Context(), reqIDKey{}, reqID))

		next.ServeHTTP(w, r)
	}
}

// mwLog logs a request, response, and timing information.
func (s *server) mwLog(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()

		// log incoming request data
		s.Log(LevelInfo, req.Context(), "<-- %s %s from %s", req.Method, req.RequestURI, req.RemoteAddr)
		rsp := &responseLogger{ResponseWriter: w}

		// pass through to next handler
		next.ServeHTTP(rsp, req)

		var level = LevelInfo
		if rsp.status >= 500 {
			level = LevelError
		} else if rsp.status >= 400 {
			level = LevelWarn
		}

		// log final request data
		s.Log(level, req.Context(), "--> %d (%s) in %s", rsp.status, http.StatusText(rsp.status), time.Since(start))
	}
}

type responseLogger struct {
	http.ResponseWriter
	status int
}

func (r *responseLogger) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
