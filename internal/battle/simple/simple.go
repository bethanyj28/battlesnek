package simple

import (
	"math/rand"

	"github.com/bethanyj28/battlesnek/internal"
	"github.com/bethanyj28/battlesnek/internal/battle/util"
)

// Snake is a simple snake implementation
type Snake struct{}

// NewSnake returns a new implementation of simple snake
func NewSnake() *Snake {
	return &Snake{}
}

// Move calculates the ideal move a simple snake should take
func (s *Snake) Move(state internal.GameState) (string, error) {
	avoidSelf := util.AvoidSelf(state.You)
	avoidWall := util.AvoidWall(state.Board, state.You.Head)
	options := []string{}

	if len(avoidWall) < 4 { // which walls to worry about
		options = findAvailable(avoidSelf, avoidWall)
	}

	return options[rand.Intn(len(options))], nil
}

// Info returns the style info for a simple snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}

func findAvailable(self, wall []string) []string {
	available := []string{}
	for _, s := range self {
		for _, w := range wall {
			if s == w {
				available = append(available, s)
				break
			}
		}
	}
	return available
}
