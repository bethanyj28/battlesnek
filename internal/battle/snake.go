package battle

import (
	"github.com/bethanyj28/battlesnek/internal"
)

// Snake defines the actions a snake should do
type Snake interface {
	Move(state internal.GameState) (string, error)
	Info() internal.Style
}
