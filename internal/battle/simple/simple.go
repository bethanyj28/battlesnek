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
func (s *Snake) Move(state internal.GameState) (internal.Action, error) {
	avoidSelf := util.AvoidSelf(state.You)
	avoidWall := util.AvoidWall(state.Board, state.You.Head)
	avoidOthers := util.AvoidOthers(state.Board, state.You.Head)
	possibleDirections := findPossible(avoidSelf, avoidWall, avoidOthers) // these are things that are v important to avoid
	if len(possibleDirections) == 0 {                                     // go out in style
		return internal.Action{Move: "up", Shout: "Like, comment, and subscribe"}, nil
	}

	food := make([]string, 4)
	switch {
	case state.You.Health > 50:
		food = util.AvoidFood(state.Board, state.You.Head)
	case state.You.Health <= 25:
		food = util.FindFood(state.Board, state.You.Head)
	}

	return internal.Action{Move: findOptimal(possibleDirections, food)}, nil
}

// Info returns the style info for a simple snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}

func findOptimal(solved ...[]string) string {
	options := map[string]int{}
	optimal := struct {
		dir   string
		count int
	}{
		dir:   "",
		count: 0,
	}

	for _, directions := range solved {
		for _, direction := range directions {
			if _, ok := options[direction]; !ok {
				options[direction] = 1
			} else {
				options[direction]++
			}

			if options[direction] > optimal.count {
				optimal.dir = direction
				optimal.count = options[direction]
			}
		}
	}

	return optimal.dir
}

func findPossible(solved ...[]string) []string {
	options := map[string]int{}

	for _, directions := range solved {
		for _, direction := range directions {
			if _, ok := options[direction]; !ok {
				options[direction] = 1
			} else {
				options[direction]++
			}
		}
	}

	best := []string{}
	for k, v := range options {
		if v == len(solved) {
			best = append(best, k)
		}
	}

	return best
}
