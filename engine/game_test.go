package engine

import (
	"testing"
)

func TestGame(*testing.T) {
	game := Game{}
	game.Display()
}

func TestNewGame(*testing.T) {
	game := NewGame()
	game.Display()
}
