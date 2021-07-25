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
func (s *Snake) Move(state internal.GameState) (string, error) {
	choice := rand.Intn(4)
	switch choice {
	case 0:
		return "left", nil
	case 1:
		return "right", nil
	case 2:
		return "up", nil
	case 3:
		return "down", nil
	}

	return "", nil
}

// Info returns style information about this snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}
