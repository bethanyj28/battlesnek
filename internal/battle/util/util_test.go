package util

import (
	"sort"
	"testing"

	"github.com/bethanyj28/battlesnek/internal"
	"github.com/matryer/is"
)

func TestAvoidSelf(t *testing.T) {
	type testcase struct {
		name     string
		input    internal.Battlesnake
		expected []string
	}

	testcases := []testcase{
		{
			name: "short snake",
			input: internal.Battlesnake{
				Head: internal.Coord{X: 0, Y: 0},
				Body: []internal.Coord{{X: 0, Y: 0}},
			},
			expected: []string{"down", "left", "right", "up"},
		},
		{
			name: "medium straight snake",
			input: internal.Battlesnake{
				Head: internal.Coord{X: 4, Y: 4},
				Body: []internal.Coord{
					{X: 4, Y: 4},
					{X: 4, Y: 3},
					{X: 4, Y: 2},
				},
			},
			expected: []string{"left", "right", "up"},
		},
		{
			name: "long curvy snake",
			input: internal.Battlesnake{
				Head: internal.Coord{X: 3, Y: 3},
				Body: []internal.Coord{
					{X: 3, Y: 3},
					{X: 3, Y: 2},
					{X: 4, Y: 2},
					{X: 4, Y: 3},
					{X: 4, Y: 4},
					{X: 3, Y: 4},
					{X: 2, Y: 4},
				},
			},
			expected: []string{"left"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := AvoidSelf(tc.input)
			sort.Strings(actual)
			is.Equal(actual, tc.expected)
		})
	}
}

func TestAvoidWall(t *testing.T) {
	type testcase struct {
		name       string
		inputBoard internal.Board
		inputCoord internal.Coord
		expected   []string
	}

	testcases := []testcase{
		{
			name:       "avoid bottom wall",
			inputBoard: internal.Board{Width: 7, Height: 7},
			inputCoord: internal.Coord{X: 3, Y: 0},
			expected:   []string{"left", "right", "up"},
		},
		{
			name:       "avoid top wall",
			inputBoard: internal.Board{Width: 7, Height: 7},
			inputCoord: internal.Coord{X: 3, Y: 6},
			expected:   []string{"down", "left", "right"},
		},
		{
			name:       "avoid left wall",
			inputBoard: internal.Board{Width: 7, Height: 7},
			inputCoord: internal.Coord{X: 0, Y: 3},
			expected:   []string{"down", "right", "up"},
		},
		{
			name:       "avoid right wall",
			inputBoard: internal.Board{Width: 7, Height: 7},
			inputCoord: internal.Coord{X: 6, Y: 3},
			expected:   []string{"down", "left", "up"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := AvoidWall(tc.inputBoard, tc.inputCoord)
			sort.Strings(actual)
			is.Equal(actual, tc.expected)
		})
	}
}

func TestPotentialPositions(t *testing.T) {
	type testcase struct {
		name     string
		input    internal.Coord
		expected choices
	}

	testcases := []testcase{
		{
			name:  "success",
			input: internal.Coord{X: 1, Y: 1},
			expected: choices{
				left:  internal.Coord{X: 0, Y: 1},
				right: internal.Coord{X: 2, Y: 1},
				up:    internal.Coord{X: 1, Y: 2},
				down:  internal.Coord{X: 1, Y: 0},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := potentialPositions(tc.input)
			is.Equal(actual, tc.expected)
		})
	}
}

func TestAvoidFood(t *testing.T) {
	type testcase struct {
		name       string
		inputBoard internal.Board
		inputCoord internal.Coord
		expected   []string
	}

	testcases := []testcase{
		{
			name: "avoid single food",
			inputBoard: internal.Board{Food: []internal.Coord{
				{X: 3, Y: 2}, // up
			}},
			inputCoord: internal.Coord{X: 3, Y: 1},
			expected:   []string{"down", "left", "right"},
		},
		{
			name: "avoid multiple foods",
			inputBoard: internal.Board{Food: []internal.Coord{
				{X: 3, Y: 2}, // up
				{X: 3, Y: 0}, // down
			}},
			inputCoord: internal.Coord{X: 3, Y: 1},
			expected:   []string{"left", "right"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := AvoidFood(tc.inputBoard, tc.inputCoord)
			sort.Strings(actual)
			is.Equal(actual, tc.expected)
		})
	}
}

func TestFindFood(t *testing.T) {
	type testcase struct {
		name       string
		inputBoard internal.Board
		inputCoord internal.Coord
		expected   []string
	}

	testcases := []testcase{
		{
			name: "find single food",
			inputBoard: internal.Board{Food: []internal.Coord{
				{X: 3, Y: 2}, // up
			}},
			inputCoord: internal.Coord{X: 3, Y: 1},
			expected:   []string{"up"},
		},
		{
			name: "find multiple foods",
			inputBoard: internal.Board{Food: []internal.Coord{
				{X: 3, Y: 2}, // up
				{X: 3, Y: 0}, // down
			}},
			inputCoord: internal.Coord{X: 3, Y: 1},
			expected:   []string{"down", "up"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := FindFood(tc.inputBoard, tc.inputCoord)
			sort.Strings(actual)
			is.Equal(actual, tc.expected)
		})
	}
}

func TestAvoidCollisions(t *testing.T) {
	type testcase struct {
		name        string
		inputSelf   internal.Battlesnake
		inputOthers []internal.Battlesnake
		expected    []string
	}

	testcases := []testcase{
		{
			name: "other snake longer",
			inputSelf: internal.Battlesnake{
				ID:   "me",
				Head: internal.Coord{X: 0, Y: 1},
				Body: []internal.Coord{
					{X: 0, Y: 1},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
				},
				Length: 3,
			},
			inputOthers: []internal.Battlesnake{
				{
					ID:   "other1",
					Head: internal.Coord{X: 1, Y: 0},
					Body: []internal.Coord{
						{X: 1, Y: 0},
						{X: 2, Y: 0},
						{X: 3, Y: 0},
						{X: 4, Y: 0},
						{X: 5, Y: 0},
					},
					Length: 5,
				},
				{
					ID:   "me",
					Head: internal.Coord{X: 0, Y: 1},
					Body: []internal.Coord{
						{X: 0, Y: 1},
						{X: 0, Y: 2},
						{X: 0, Y: 3},
					},
					Length: 3,
				},
			},
			expected: []string{"left", "up"},
		},
		{
			name: "other snake shorter",
			inputSelf: internal.Battlesnake{
				ID:   "me",
				Head: internal.Coord{X: 0, Y: 1},
				Body: []internal.Coord{
					{X: 0, Y: 0},
					{X: 0, Y: 1},
					{X: 0, Y: 2},
				},
				Length: 3,
			},
			inputOthers: []internal.Battlesnake{
				{
					ID:   "other1",
					Head: internal.Coord{X: 1, Y: 0},
					Body: []internal.Coord{
						{X: 1, Y: 0},
						{X: 2, Y: 0},
					},
					Length: 2,
				},
				{
					ID:   "me",
					Head: internal.Coord{X: 0, Y: 1},
					Body: []internal.Coord{
						{X: 0, Y: 0},
						{X: 0, Y: 1},
						{X: 0, Y: 2},
					},
					Length: 3,
				},
			},
			expected: []string{"down", "left", "right", "up"},
		},
		{
			name: "multiple snakes",
			inputSelf: internal.Battlesnake{
				ID:   "me",
				Head: internal.Coord{X: 3, Y: 1},
				Body: []internal.Coord{
					{X: 3, Y: 3},
					{X: 3, Y: 2},
					{X: 3, Y: 1},
				},
				Length: 3,
			},
			inputOthers: []internal.Battlesnake{
				{
					ID:   "other1",
					Head: internal.Coord{X: 2, Y: 0},
					Body: []internal.Coord{
						{X: 2, Y: 0},
						{X: 1, Y: 0},
						{X: 0, Y: 0},
					},
					Length: 3,
				},
				{
					ID:   "other2",
					Head: internal.Coord{X: 5, Y: 1},
					Body: []internal.Coord{
						{X: 5, Y: 1},
						{X: 6, Y: 1},
						{X: 7, Y: 1},
					},
					Length: 3,
				},
				{
					ID:   "me",
					Head: internal.Coord{X: 3, Y: 1},
					Body: []internal.Coord{
						{X: 3, Y: 3},
						{X: 3, Y: 2},
						{X: 3, Y: 1},
					},
					Length: 3,
				},
			},
			expected: []string{"up"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := AvoidCollisions(tc.inputSelf, tc.inputOthers)
			sort.Strings(actual)
			is.Equal(actual, tc.expected)
		})
	}
}

// 03 13 23 33 43
// 02 12 22 32 42
// 01 11 21 31 41
// 00 10 20 30 40

func TestIntrovertSnake(t *testing.T) {
	type testcase struct {
		name        string
		inputSelf   internal.Battlesnake
		inputOthers []internal.Battlesnake
		expected    []string
	}

	testcases := []testcase{
		{
			name: "other snake longer",
			inputSelf: internal.Battlesnake{
				ID:   "me",
				Head: internal.Coord{X: 0, Y: 1},
				Body: []internal.Coord{
					{X: 0, Y: 1},
					{X: 0, Y: 2},
					{X: 0, Y: 3},
				},
				Length: 3,
			},
			inputOthers: []internal.Battlesnake{
				{
					ID:   "other1",
					Head: internal.Coord{X: 3, Y: 2},
					Body: []internal.Coord{
						{X: 3, Y: 2},
						{X: 2, Y: 2},
						{X: 2, Y: 3},
						{X: 3, Y: 3},
						{X: 4, Y: 3},
					},
					Length: 5,
				},
				{
					ID:   "me",
					Head: internal.Coord{X: 0, Y: 1},
					Body: []internal.Coord{
						{X: 0, Y: 1},
						{X: 0, Y: 2},
						{X: 0, Y: 3},
					},
					Length: 3,
				},
			},
			expected: []string{"down", "left"},
		},
		{
			name: "multiple snakes",
			inputSelf: internal.Battlesnake{
				ID:   "me",
				Head: internal.Coord{X: 0, Y: 0},
				Body: []internal.Coord{
					{X: 0, Y: 0},
				},
				Length: 1,
			},
			inputOthers: []internal.Battlesnake{
				{
					ID:   "other1",
					Head: internal.Coord{X: 5, Y: 5},
					Body: []internal.Coord{
						{X: 5, Y: 5},
					},
					Length: 1,
				},
				{
					ID:   "other2",
					Head: internal.Coord{X: 5, Y: 0},
					Body: []internal.Coord{
						{X: 5, Y: 0},
					},
					Length: 1,
				},
				{
					ID:   "me",
					Head: internal.Coord{X: 0, Y: 0},
					Body: []internal.Coord{
						{X: 0, Y: 0},
					},
					Length: 1,
				},
			},
			expected: []string{"down", "left", "right", "up"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)
			actual := IntrovertSnake(tc.inputSelf, tc.inputOthers, 5)
			sort.Strings(actual)

			is.Equal(actual, tc.expected)
		})
	}
}
