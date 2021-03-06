package clairvoyant

import "github.com/bethanyj28/battlesnek/internal"

// Snake is a clairvoyant implementation of snake
type Snake struct {
	lookahead int
}

// NewSnake creates a new clairvoyant snake
func NewSnake(lookahead int) *Snake {
	if lookahead == 0 {
		lookahead = 1
	}
	return &Snake{lookahead: lookahead}
}

// Move decides which move is ideal based on seeing the future
func (s *Snake) Move(state internal.GameState) (internal.Action, error) {
	return internal.Action{Move: findOptimal(state, s.lookahead)}, nil
}

// Info ensures the snake is stylin and profilin
func (s *Snake) Info() internal.Style {
	return internal.Style{
		//Color: "#76A5AF",
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}
