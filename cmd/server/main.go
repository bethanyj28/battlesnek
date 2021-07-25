package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config represents environment variables
type Config struct {
	Host    string        `default:"0.0.0.0"`
	Port    string        `envconfig:"PORT" default:"8080"`
	Timeout time.Duration `default:"5s"`
}

func main() {
	logger := logrus.New()
	var c Config
	if err := envconfig.Process("api", &c); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to read envconfig").Error())
	}

	svr := newServer(fmt.Sprintf("%s:%s", c.Host, c.Port), c.Timeout, logger)

	log.Fatal(svr.httpServer.ListenAndServe())
}
