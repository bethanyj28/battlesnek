package main

import (
	"net/http"
	"time"

	"github.com/bethanyj28/battlesnek/internal/battle"
	"github.com/bethanyj28/battlesnek/internal/battle/clairvoyant"
	"github.com/sirupsen/logrus"
)

type server struct {
	httpServer *http.Server
	logger     *logrus.Logger
	snake      battle.Snake
}

func newServer(addr string, timeout time.Duration, logger *logrus.Logger) *server {
	svr := &server{logger: logger}
	svr.buildHTTPServer(addr, timeout)

	svr.snake = clairvoyant.NewSnake(5)

	return svr
}
