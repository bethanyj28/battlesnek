package internal

import "errors"

// ErrUnimplemented is returned when a method has not been fleshed out yet
var ErrUnimplemented error = errors.New("method unimplemented")
