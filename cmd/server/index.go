package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func (s *server) handleIndex() http.HandlerFunc {
	type response struct {
		APIVersion string `json:"apiversion"`
		Author     string `json:"author"`
		Color      string `json:"color"`
		Head       string `json:"head"`
		Tail       string `json:"tail"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		info := s.snake.Info()
		resp := response{
			APIVersion: "1",
			Author:     "bethanyj28",
			Color:      info.Color,
			Head:       info.Head,
			Tail:       info.Tail,
		}
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}).Error("failed to marshal info")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
