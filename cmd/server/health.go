package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *server) handleHealth() http.HandlerFunc {
	type response struct {
		Healthy bool `json:"healthy"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(response{Healthy: true})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}).Error("failed to check health")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
