package main

import (
	"log"

	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/patterns"
)

func main() {

	ui := d2dui.New()
	matrix := game.NewMatrix(60, 60)
	patterns.Glider(matrix, 1, 1)
	patterns.Pulsar(matrix, 24, 24)

	game := game.New(ui, game.Options{
		Fps:           1,
		InitialMatrix: matrix,
	})
	ui.SetGame(game)
	game.Play()

	err := ui.Start()
	if err != nil {
		log.Println(err)
	}
}
