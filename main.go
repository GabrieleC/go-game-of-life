package main

import (
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/game/conwaylogic"
	"gcoletta.it/game-of-life/internal/patterns"
)

func main() {

	ui := d2dui.New(800, 800)

	matrix := game.NewMatrix(60, 60)
	patterns.Glider(matrix, 1, 1)
	patterns.Pulsar(matrix, 24, 24)

	game := game.New(ui, conwaylogic.Iterate,
		game.Options{
			Fps:           2,
			InitialMatrix: matrix,
		})
	game.Play()
	game.Start()
}
