package util

import (
	"fmt"
	"math"

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

// AvoidHazards returns moves that prevent snake from running into hazards
func AvoidHazards(board internal.Board, head internal.Coord) []string {
	pos := potentialPositions(head)
	avoid := convertCoordsToGrid(board.Hazards)

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

// MoveAwayFromSelf returns moves that are further from the snake's center of self
func MoveAwayFromSelf(self internal.Battlesnake) map[string]int {
	pos := potentialPositions(self.Head)
	avgPos := averagePositions(self.Body)

	distMap := map[string]int{}

	for dir, p := range pos {
		xDiffSquare := math.Pow(float64(p.X-avgPos.X), 2)
		yDiffSquare := math.Pow(float64(p.Y-avgPos.Y), 2)
		dist := math.Sqrt(xDiffSquare + yDiffSquare)
		distMap[dir.String()] = int(math.Round(dist))
	}

	return distMap
}

// AvoidCollisions returns moves that avoid collisions with stronger snakes
func AvoidCollisions(self internal.Battlesnake, others []internal.Battlesnake) []string {
	otherPosCoords := []internal.Coord{}
	for _, snake := range others {
		if snake.ID == self.ID {
			continue
		}

		if snake.Length < self.Length {
			continue
		}

		otherPosCoords = append(otherPosCoords, potentialPositionsSlice(snake.Head)...)
	}

	fmt.Println(otherPosCoords)

	return decideDir(potentialPositions(self.Head), convertCoordsToGrid(otherPosCoords))
}

func averagePositions(coords []internal.Coord) internal.Coord {
	sumX := 0
	sumY := 0

	for _, coord := range coords {
		sumX += coord.X
		sumY += coord.Y
	}

	avgX := sumX / len(coords)
	avgY := sumY / len(coords)

	return internal.Coord{X: avgX, Y: avgY}
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

func potentialPositionsSlice(head internal.Coord) []internal.Coord {
	positions := []internal.Coord{}
	for _, pos := range potentialPositions(head) {
		positions = append(positions, pos)
	}

	return positions
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
