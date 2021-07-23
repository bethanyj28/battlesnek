package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	r.HandleFunc("/health", s.handleHealth())

	return r
}
