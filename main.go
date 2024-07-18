package main

import (
	"gcoletta.it/game-of-life/internal/gui"
	"gcoletta.it/game-of-life/internal/game"
)

func main() {

	g := gui.New(30, 30, gui.Options{CellSize: 20})
	game := game.New(g, 1)
	game.Execute()

	err := g.Start()
	if err != nil {
		panic(err)
	}
}
