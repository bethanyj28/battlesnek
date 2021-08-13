package util

import (
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

// Simplifies ugly types
type matrix map[int]map[int]bool
type choices map[direction]internal.Coord

// AvoidSelf returns the moves that will prevent the snake from running into itself
func AvoidSelf(self internal.Battlesnake) []string {
	pos := potentialPositions(self.Head)
	avoid := convertCoordsToGrid(self.Body)

	return decideDir(pos, avoid)
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

// AvoidOthers returns moves that will prevent the snake from running into others
func AvoidOthers(board internal.Board, head internal.Coord) []string {
	pos := potentialPositions(head)
	enemyPos := []internal.Coord{}
	for _, enemy := range board.Snakes {
		enemyPos = append(enemyPos, enemy.Body...)
		enemyPos = append(enemyPos, enemy.Head)
	}

	avoid := convertCoordsToGrid(enemyPos)

	return decideDir(pos, avoid)
}

// FindFood returns moves that gives food
func FindFood(board internal.Board, head internal.Coord) []string {
	noFood := AvoidFood(board, head)
	return inverse(noFood)
}

// AvoidFood returns moves that avoids food
func AvoidFood(board internal.Board, head internal.Coord) []string {
	pos := potentialPositions(head)
	avoid := convertCoordsToGrid(board.Food)

	return decideDir(pos, avoid)
}

// TODO: cache these
func potentialPositions(head internal.Coord) choices {
	return choices{
		left:  internal.Coord{X: head.X - 1, Y: head.Y},
		right: internal.Coord{X: head.X + 1, Y: head.Y},
		up:    internal.Coord{X: head.X, Y: head.Y + 1},
		down:  internal.Coord{X: head.X, Y: head.Y - 1},
	}
}

func convertCoordsToGrid(coords []internal.Coord) matrix {
	grid := matrix{}
	for _, coord := range coords {
		if _, ok := grid[coord.X]; !ok {
			grid[coord.X] = map[int]bool{}
		}
		grid[coord.X][coord.Y] = true
	}

	return grid
}

func decideDir(potential choices, avoid matrix) []string {
	options := []string{}

	for dir, coord := range potential {
		if _, ok := avoid[coord.X][coord.Y]; ok {
			continue
		}
		options = append(options, dir.String())
	}

	return options
}

func inverse(directions []string) []string {
	dirMap := map[string]bool{
		left.String():  true,
		right.String(): true,
		up.String():    true,
		down.String():  true,
	}

	for _, dir := range directions {
		dirMap[dir] = false
	}

	invertedDirections := []string{}
	for k, v := range dirMap {
		if v {
			invertedDirections = append(invertedDirections, k)
		}
	}

	return invertedDirections
}
