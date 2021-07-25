package main

import (
	"encoding/json"
	"net/http"

	"github.com/bethanyj28/battlesnek/internal"
	"github.com/sirupsen/logrus"
)

func (s *server) handleEnd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := internal.GameState{}
		if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}).Error("failed to decode game state on end")
			return
		}

		s.logger.WithFields(logrus.Fields{
			"game_state": state,
		}).Info("END")

		w.WriteHeader(http.StatusOK)
	}
}
