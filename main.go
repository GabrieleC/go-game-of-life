package main

import (
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/game/conwaylogic"
)

func main() {

	ui := d2dui.New(840, 840)

	matrix := game.NewMatrix(60, 60)

	game := game.New(ui, conwaylogic.Iterate,
		game.Options{
			Fps:           2,
			InitialMatrix: matrix,
		})
	game.Play()
	game.Start()
}
