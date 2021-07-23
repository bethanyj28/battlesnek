package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config represents environment variables
type Config struct {
	Address string        `default:"0.0.0.0:8080"`
	Timeout time.Duration `default:"5s"`
}

func main() {
	logger := logrus.New()
	var c Config
	if err := envconfig.Process("api", &c); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to read envconfig").Error())
	}

	svr := newServer(c.Address, c.Timeout, logger)

	log.Fatal(svr.httpServer.ListenAndServe())
}
