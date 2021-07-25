package util

import "github.com/bethanyj28/battlesnek/internal"

type direction int

const (
	left direction = iota + 1
	right
	up
	down
)

func (d direction) String() string {
	return [...]string{"left", "right", "up", "down"}[d]
}

// AvoidSelf returns the moves that will prevent the snake from running into itself
func AvoidSelf(self internal.Battlesnake) []string {
	pos := potentialPositions(self.Head)
	avoid := convertCoordsToGrid(self.Body)
	options := []string{}

	for dir, coord := range pos {
		if _, ok := avoid[coord.X][coord.Y]; !ok {
			continue
		}
		options = append(options, dir.String())
	}

	return options
}

// AvoidWall returns moves that will prevent the snake from running into a wall
func AvoidWall(board internal.Board, head internal.Coord) []string {
	pos := potentialPositions(head)
	options := []string{}

	for dir, coord := range pos {
		if coord.X >= board.Width || coord.X < 0 {
			continue
		}

		if coord.Y >= board.Height || coord.Y < 0 {
			continue
		}

		options = append(options, dir.String())
	}

	return options
}

func potentialPositions(head internal.Coord) map[direction]internal.Coord {
	return map[direction]internal.Coord{
		left:  internal.Coord{X: head.X - 1, Y: head.Y},
		right: internal.Coord{X: head.X + 1, Y: head.Y},
		up:    internal.Coord{X: head.X, Y: head.Y + 1},
		down:  internal.Coord{X: head.X, Y: head.Y - 1},
	}
}

func convertCoordsToGrid(coords []internal.Coord) map[int]map[int]bool {
	grid := map[int]map[int]bool{}
	for _, coord := range coords {
		if _, ok := grid[coord.X]; !ok {
			grid[coord.X] = map[int]bool{}
		}
		grid[coord.X][coord.Y] = true
	}

	return grid
}
