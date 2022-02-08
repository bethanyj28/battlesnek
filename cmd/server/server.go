package main

import (
	"net/http"
	"time"

	"github.com/bethanyj28/battlesnek/internal/battle"
	"github.com/bethanyj28/battlesnek/internal/battle/clairvoyant"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type server struct {
	httpServer *http.Server
	logger     *logrus.Logger
	snake      battle.Snake
	nrApp      *newrelic.Application
}

func newServer(addr string, timeout time.Duration, nr *newrelic.Application, logger *logrus.Logger) *server {
	svr := &server{logger: logger, nrApp: nr}
	svr.buildHTTPServer(addr, timeout)

	svr.snake = clairvoyant.NewSnake(5)

	return svr
}
