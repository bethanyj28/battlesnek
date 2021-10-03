package clairvoyant

import "github.com/bethanyj28/battlesnek/internal"

// Snake is a clairvoyant implementation of snake
type Snake struct {
	lookahead int
}

// NewClairvoyantSnake creates a new clairvoyant snake
func NewClairvoyantSnake(lookahead int) Snake {
	if lookahead == 0 {
		lookahead = 1
	}
	return Snake{lookahead: lookahead}
}

// Move decides which move is ideal based on seeing the future
func (s *Snake) Move(state internal.GameState) (internal.Action, error) {
	moves := drillDown(state, s.lookahead)

	return internal.Action{Move: moves.move}, nil
}

// Info ensures the snake is stylin and profilin
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}

type route struct {
	move   string
	food   int
	hazard int
}

func drillDown(state internal.GameState, lookahead int) route {
	return route{}
}
