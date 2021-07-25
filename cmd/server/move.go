package main

import (
	"encoding/json"
	"net/http"

	"github.com/bethanyj28/battlesnek/internal"
	"github.com/sirupsen/logrus"
)

func (s *server) handleMove() http.HandlerFunc {
	type response struct {
		Move  string `json:"move"`
		Shout string `json:"shout,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		state := internal.GameState{}
		if err := json.NewDecoder(r.Body).Decode(&state); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusBadRequest,
				"error": err.Error(),
			}).Error("failed to decode game state on move")
			return

		}

		move, err := s.snake.Move(state)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}).Error("failed to choose move")
			return
		}

		resp := response{
			Move: move,
		}

		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			s.logger.WithFields(logrus.Fields{
				"code":  http.StatusInternalServerError,
				"error": err.Error(),
			}).Error("failed to marshal move")
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
