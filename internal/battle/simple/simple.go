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
	switch len(possibleDirections) {
	case 0: // go out in style
		return internal.Action{Move: "up", Shout: "Like, comment, and subscribe"}, nil
	case 1: // don't do extra work
		return internal.Action{Move: possibleDirections[0]}, nil
	}

	// Calculate need to find food and prioritize
	food := make([]string, 4)
	switch {
	case state.You.Health > 75:
		food = util.AvoidFood(state.Board, state.You.Head)
	case state.You.Health <= 25:
		food = util.FindFood(state.Board, state.You.Head)
	}

	foodPriorityMap := map[string]int{}
	for _, dir := range food {
		foodPriorityMap[dir] = 1
	}

	// Avoid hazards if possible
	avoidHazards := util.AvoidHazards(state.Board, state.You.Head)
	avoidHazardsPriorityMap := map[string]int{}
	healthMultiplier := 2
	switch {
	case state.You.Health > 75:
		healthMultiplier = 1
	case state.You.Health < 50:
		healthMultiplier = 4
	}
	for _, dir := range avoidHazards {
		avoidHazardsPriorityMap[dir] = healthMultiplier
	}

	// avoid collisions with stronk snakes
	avoidCollisions := util.AvoidCollisions(state.You, state.Board.Snakes)
	avoidCollisionsPriorityMap := map[string]int{}
	for _, dir := range avoidCollisions {
		avoidCollisionsPriorityMap[dir] = 3
	}

	avoidSelfPriorityMap := util.MoveAwayFromSelf(state.You)
	introvertPriorityMap := util.IntrovertSnake(state.You, state.Board.Snakes)

	return internal.Action{Move: findOptimal(possibleDirections, foodPriorityMap, avoidSelfPriorityMap, avoidHazardsPriorityMap, avoidCollisionsPriorityMap, introvertPriorityMap)}, nil
}

// Info returns the style info for a simple snake
func (s *Snake) Info() internal.Style {
	return internal.Style{
		Color: "#00cc99",
		Head:  "beluga",
		Tail:  "freckled",
	}
}

func findOptimal(available []string, prioritized ...map[string]int) string {
	options := map[string]int{}
	optimal := struct {
		dir   string
		count int
	}{
		dir:   "",
		count: 0,
	}

	for _, a := range available {
		options[a] = 1
		// initialize optimal with something
		optimal.dir = a
		optimal.count = 1
	}

	for _, directions := range prioritized {
		for k, v := range directions {
			if _, ok := options[k]; !ok { // don't go that direction
				continue
			} else {
				options[k] += v
			}

			if options[k] > optimal.count {
				optimal.dir = k
				optimal.count = options[k]
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
