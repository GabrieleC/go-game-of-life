package main

import (
	"log"

	"gcoletta.it/game-of-life/internal/ebidenUi"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/patterns"
)

func main() {

	gui := ebidenUi.New(60, 120, ebidenUi.Options{CellSize: 10})
	matrix := game.NewMatrix(60, 120)
	patterns.Glider(game.Matrix(matrix), 1, 1)
	patterns.Pulsar(game.Matrix(matrix), 24, 24)

	game := game.New(gui, game.Options{
		Fps:           1,
		InitialMatrix: matrix,
	})
	game.Execute()

	err := gui.Start()
	if err != nil {
		log.Println(err)
	}
}
