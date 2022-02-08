package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config represents environment variables
type Config struct {
	Host               string        `default:"0.0.0.0"`
	Port               string        `envconfig:"PORT" default:"8080"`
	Timeout            time.Duration `default:"5s"`
	NewRelicName       string
	NewRelicLicenseKey string
}

func main() {
	logger := logrus.New()
	var c Config
	if err := envconfig.Process("api", &c); err != nil {
		logger.Fatal(errors.Wrap(err, "failed to read envconfig").Error())
	}

	app, err := newrelic.NewApplication(createNewRelicConfigOptions(c.NewRelicName, c.NewRelicLicenseKey)...)
	if err != nil {
		logger.Fatal(errors.Wrap(err, "failed to configure new relic").Error())
	}

	svr := newServer(fmt.Sprintf("%s:%s", c.Host, c.Port), c.Timeout, app, logger)

	log.Fatal(svr.httpServer.ListenAndServe())
}

func createNewRelicConfigOptions(name, key string) []newrelic.ConfigOption {
	if name == "" || key == "" {
		return []newrelic.ConfigOption{newrelic.ConfigEnabled(false)}
	}
	return []newrelic.ConfigOption{
		newrelic.ConfigAppName(name),
		newrelic.ConfigLicense(key),
		newrelic.ConfigDistributedTracerEnabled(true),
	}
}
