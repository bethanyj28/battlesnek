package simple

import (
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
	options := append(avoidSelf, avoidWall...)

	return findOptimal(options), nil
}

// Info returns the style info for a simple snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}

func findOptimal(options []string) string {
	counts := map[string]int{}
	for _, opts := range options {
		if _, ok := counts[opts]; !ok {
			counts[opts] = 1
			continue
		}
		counts[opts]++
	}

	max := 0
	optimal := ""
	for k, v := range counts {
		if v > max {
			optimal = k
		}
	}

	return optimal
}
