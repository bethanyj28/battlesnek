package internal

// GameState describes the state of the board
type GameState struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

// Game describes the game being played
type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Timeout int32   `json:"timeout"`
}

// Ruleset describes the rules of the game
type Ruleset struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Board describes characteristics of the board at a given point
type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`

	// Used in non-standard game modes
	Hazards []Coord `json:"hazards"`
}

// Battlesnake describes properties of your snake at the moment
type Battlesnake struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Health  int32   `json:"health"`
	Body    []Coord `json:"body"`
	Head    Coord   `json:"head"`
	Length  int32   `json:"length"`
	Latency string  `json:"latency"`

	// Used in non-standard game modes
	Shout string `json:"shout"`
	Squad string `json:"squad"`
}

// Coord are the x, y position of the snake on the board
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Equals returns if the coordinate is equal to another one
func (c Coord) Equals(other Coord) bool {
	if c.X == other.X && c.Y == other.Y {
		return true
	}
	return false
}

// Style is the style of the snake
type Style struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

// Action is the action a snake takes on move
type Action struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}
