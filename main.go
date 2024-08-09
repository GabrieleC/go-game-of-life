package main

import (
	"gcoletta.it/game-of-life/internal/conwaylogic"
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
)

func main() {

	ui := d2dui.New(840, 840)
	matrix := game.NewMatrix(60, 60)
	game := game.New(ui, conwaylogic.Iterate,
		game.Options{
			Fps:           2,
			InitialMatrix: matrix,
		})
		
	game.Start()
}
