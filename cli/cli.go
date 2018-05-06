package cli

import (
	"fmt"

	"github.com/mashumafi/gladiator/engine"
)

func Process() {
	game := engine.NewGame()
	fmt.Println(game.Display())
}
