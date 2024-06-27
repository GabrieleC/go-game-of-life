package main

import (
	"log"

	"gcoletta.it/2d2/gui"
)

func main() {

	cellSize := float32(20)
	gridSize := 30

	g, matrix := gui.Make(gridSize, gui.MakeOptions{CellSize: &cellSize})

	matrix[5][0] = true

	if err := g.Start(); err != nil {
		log.Fatal(err)
	}

}
