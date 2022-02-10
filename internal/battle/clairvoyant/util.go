package clairvoyant

import (
	"log"

	"github.com/bethanyj28/battlesnek/internal"
)

type direction int

const (
	left direction = iota + 1
	right
	up
	down
)

// Convert direction to string
func (d direction) String() string {
	return [...]string{"", "left", "right", "up", "down"}[d]
}

type boardObject int

const (
	food boardObject = iota + 1
	otherSnake
	you
	hazard
)

func (b boardObject) String() string {
	return [...]string{"", "F", "O", "Y", "H"}[b]
}

// Simplifies ugly types
type matrix map[int]map[int][]string
type choices map[direction]internal.Coord

func convertBoardToMatrix(state internal.GameState) matrix {
	board := state.Board
	grid := matrix{}
	for _, f := range board.Food {
		grid = addCoordToMatrix(f, grid, food.String())
	}

	for _, snake := range board.Snakes {
		obj := otherSnake
		if snake.ID == state.You.ID {
			obj = you
		}
		for _, coord := range snake.Body {
			grid = addCoordToMatrix(coord, grid, obj.String())
		}
	}

	for _, h := range board.Hazards {
		grid = addCoordToMatrix(h, grid, hazard.String())
	}

	return grid
}

func addCoordToMatrix(coord internal.Coord, grid matrix, object string) matrix {
	if _, ok := grid[coord.X]; !ok {
		grid[coord.X] = map[int][]string{}
	}

	if _, ok := grid[coord.X][coord.Y]; !ok {
		grid[coord.X][coord.Y] = []string{object}
		return grid
	}

	grid[coord.X][coord.Y] = append(grid[coord.X][coord.Y], object)
	return grid
}

func potentialPositions(head internal.Coord) choices {
	return choices{
		left:  internal.Coord{X: head.X - 1, Y: head.Y},
		right: internal.Coord{X: head.X + 1, Y: head.Y},
		up:    internal.Coord{X: head.X, Y: head.Y + 1},
		down:  internal.Coord{X: head.X, Y: head.Y - 1},
	}
}

func checkPossible(initialState internal.GameState, grid matrix, potential internal.Coord, otherSnakeLength map[string]int32) bool {
	// check walls
	if initialState.Game.Ruleset.Name != "wrapped" {
		if potential.X >= initialState.Board.Width || potential.X < 0 {
			return false
		}

		if potential.Y >= initialState.Board.Height || potential.Y < 0 {
			return false
		}
	}

	// check hazards
	space := grid[potential.X][potential.Y]
	for _, obj := range space {
		if obj == you.String() || obj == otherSnake.String() {
			return false
		}

		if l, ok := otherSnakeLength[obj]; ok {
			if l > initialState.You.Length {
				return false
			}
		}
	}

	return true
}

func moveEnemiesForward(oldHead, newHead internal.Coord, snake internal.Battlesnake, grid matrix) matrix {
	objList := grid[oldHead.X][oldHead.Y]
	objList = append(objList, otherSnake.String())
	for i, obj := range objList {
		if obj == snake.ID {
			grid[oldHead.X][oldHead.Y] = append(objList[:i], objList[i+1:]...)
			break
		}
	}

	return addCoordToMatrix(newHead, grid, snake.ID)
}

func mapSnakeLengths(snakes []internal.Battlesnake) map[string]int32 {
	snakeLengths := map[string]int32{}
	for _, snake := range snakes {
		snakeLengths[snake.ID] = snake.Length
	}

	return snakeLengths
}

type route struct {
	numEval   int
	numDeaths int
	numHazard int
	numFood   int
}

func findOptimal(state internal.GameState, movesLeft int) string {
	grid := convertBoardToMatrix(state)
	for _, snake := range state.Board.Snakes {
		if snake.ID == state.You.ID {
			continue
		}
		p := potentialPositions(snake.Head)
		for _, pp := range p {
			grid = moveEnemiesForward(snake.Head, pp, snake, grid)
		}
	}

	health := float64(state.You.Health) / 100
	potential := potentialPositions(state.You.Head)
	analysis := map[direction]float64{}
	otherSnakes := mapSnakeLengths(state.Board.Snakes)
	for d, p := range potential {
		r := lookahead(movesLeft, route{}, state, otherSnakes, grid, p)
		if r.numDeaths == r.numEval {
			continue
		}
		survival := float64(r.numDeaths) / (float64(r.numEval * r.numEval))
		hazardRisk := ((1 - health) * float64(r.numHazard)) / float64(r.numEval*r.numEval)
		foodRisk := (health * float64(r.numEval-r.numFood)) / float64(r.numEval*r.numEval)
		sum := survival
		denom := 1

		sum += hazardRisk
		denom++

		sum += foodRisk
		denom++

		analysis[d] = sum / float64(denom)
		log.Printf("route for %s:%+v", d.String(), r)
	}

	log.Print(analysis)

	minScore := 1.0
	optimal := "up"
	for d, s := range analysis {
		if s < minScore {
			optimal = d.String()
			minScore = s
		}
	}
	return optimal
}

func lookahead(movesLeft int, r route, initialState internal.GameState, otherSnakes map[string]int32, grid matrix, option internal.Coord) route {
	if movesLeft <= 0 {
		return r
	}

	r.numEval++

	if !checkPossible(initialState, grid, option, otherSnakes) {
		r.numDeaths++
		return r
	}

	if obj, ok := grid[option.X][option.Y]; ok {
		for _, o := range obj {
			if o == food.String() {
				r.numFood++
			}

			if o == hazard.String() {
				r.numHazard++
			}
		}
	}

	potential := potentialPositions(option)
	grid = addCoordToMatrix(option, grid, you.String())
	for _, p := range potential {
		movesLeft--
		nr := lookahead(movesLeft, r, initialState, otherSnakes, grid, p)
		r.numEval += nr.numEval
		r.numDeaths += nr.numDeaths
		r.numHazard += nr.numHazard
		r.numFood += nr.numFood
	}

	return r
}
