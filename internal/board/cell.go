package board

import (
	"fmt"
)

// do not create interfaces just to have interfaces.
// struct is better if you don't expect it have many different implementations.

// It's enough to have this
type Cell rune

// no need for method to just read public field.
// It's idiomatic in go to not add "get" of "Get" to getter names.

// it's idiomatic to use short name for receiver arg
func (c *Cell) setValue(v rune) bool {
	if *c == 'X' || *c == 'O' {
		return false
	}

	// consider flat code over nested
	*c = value

	return true
}
