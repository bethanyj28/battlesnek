package main

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type server struct {
	httpServer *http.Server
	logger     *logrus.Logger
}

func newServer(addr string, timeout time.Duration, logger *logrus.Logger) *server {
	svr := &server{logger: logger}
	svr.buildHTTPServer(addr, timeout)

	return svr
}
