package engine

import (
	"github.com/mashumafi/gladiator/core"
	"github.com/mashumafi/gladiator/places"
)

// Game for gladiator
type Game struct {
	place  places.Place
	player core.Player
}

// NewGame setup game
func NewGame() *Game {
	game := &Game{}
	game.place = places.Town{}
	game.player = core.Player{}
	return game
}

// Display get place display
func (g *Game) Display() string {
	if g.place == nil {
		return ""
	}
	return g.place.Display()
}
