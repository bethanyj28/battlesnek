package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bethanyj28/battlesnek/internal"
	"github.com/newrelic/go-agent/v3/newrelic"
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

		if txn := newrelic.FromContext(r.Context()); nil != txn {
			txn.AddAttribute("won", didSnakeWin(state.Board.Snakes, state.You))
			txn.AddAttribute("turns", state.Turn)
			txn.AddAttribute("game_id", state.Game.ID)
		}

		s.logger.WithFields(logrus.Fields{
			"game_state": fmt.Sprintf("%+v", state),
		}).Info("END")

		w.WriteHeader(http.StatusOK)
	}
}

func didSnakeWin(remainingSnakes []internal.Battlesnake, you internal.Battlesnake) bool {
	for _, snake := range remainingSnakes {
		if snake.ID == you.ID {
			return true
		}
	}
	return false
}
