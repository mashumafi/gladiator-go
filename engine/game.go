package engine

import (
	"github.com/mashumafi/gladiator/places"
)

// Game for gladiator
type Game struct {
	place  places.Place
	player Player
}

func NewGame() *Game {
	game := &Game{}
	game.place = places.Town{}
	game.player = Player{}
	return game
}

// Play entry point
func (g Game) Play() {
}
