package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
)

func (s *server) buildHTTPServer(addr string, timeout time.Duration) {
	s.httpServer = &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * timeout,
		ReadTimeout:  time.Second * timeout,
		Handler:      s.buildRoutes(),
	}
}

func (s *server) buildRoutes() *mux.Router {
	r := mux.NewRouter()
	r.Use(nrgorilla.Middleware(s.nrApp))

	r.HandleFunc("/", s.handleIndex())
	r.HandleFunc("/health", s.handleHealth())
	r.HandleFunc("/start", s.handleStart())
	r.HandleFunc("/move", s.handleMove())
	r.HandleFunc("/end", s.handleEnd())

	return r
}
