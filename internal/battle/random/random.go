package random

import (
	"math/rand"

	"github.com/bethanyj28/battlesnek/internal"
)

// Snake is a snake that chooses a random direction as a move
type Snake struct{}

// NewSnake returns a new instance of a randomsnake
func NewSnake() *Snake {
	return &Snake{}
}

// Move chooses a random direction for the snake to go
func (s *Snake) Move(state internal.GameState) (internal.Action, error) {
	action := internal.Action{}
	choice := rand.Intn(4)
	switch choice {
	case 0:
		action.Move = "left"
	case 1:
		action.Move = "right"
	case 2:
		action.Move = "up"
	case 3:
		action.Move = "down"
	default:
		action.Move = ""
	}

	return action, nil
}

// Info returns style information about this snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}
