package main

import (
	"gcoletta.it/game-of-life/internal/conwaylogic"
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
)

func main() {

	ui := d2dui.New(840, 840)
	matrix := game.NewMatrix(120, 120)
	game := game.New(ui, conwaylogic.Iterate,
		game.Options{
			Fps:           4,
			InitialMatrix: matrix,
		})
		
	game.Start()
}
