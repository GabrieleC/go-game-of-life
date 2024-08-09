package main

import (
	"os"
	"strconv"

	"gcoletta.it/game-of-life/internal/conwaylogic"
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
)

func main() {

	size := parseSizeParameter(60)

	ui := d2dui.New(840, 840)
	matrix := game.NewMatrix(size, size)
	game := game.New(ui, conwaylogic.Iterate,
		game.Options{
			Fps:           4,
			InitialMatrix: matrix,
		})

	game.Start()
}

func parseSizeParameter(def int) int {
	if len(os.Args) > 1 {
		parsedArg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			return parsedArg
		}
	}
	return def
}
