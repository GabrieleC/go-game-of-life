package main

import (
	"fmt"
	"os"
	"strconv"

	"gcoletta.it/game-of-life/internal/conwayadapter"
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/d2duiadapter"
	"gcoletta.it/game-of-life/internal/game"
)

func main() {

	size, err := parseSizeParameter(500)
	if err != nil {
		printUsage()
		os.Exit(1)
	}

	ui := d2duiadapter.D2duiAdapter(d2dui.New(840, 840))
	mtx := game.NewMatrix(size, size)

	game := game.New(ui, conwayadapter.Iterate,
		game.Options{
			Fps:           4,
			InitialMatrix: mtx,
		})

	game.Start()
}

func parseSizeParameter(def int) (int, error) {
	if len(os.Args) > 1 {
		parsedArg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			return parsedArg, nil
		} else {
			return 0, err
		}
	}
	return def, nil
}

func printUsage() {
	fmt.Println("First parameter must be a positive integer number, defining the game universe size")
}
