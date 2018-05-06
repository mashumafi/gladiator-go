package engine

import (
	"testing"
)

func TestGame(*testing.T) {
	game := Game{}
	game.Play()
}

func TestNewGame(*testing.T) {
	game := NewGame()
	game.Play()
}
