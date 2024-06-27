package main

import (
	"log"
	"math/rand/v2"

	"gcoletta.it/2d2/gui"
)

func main() {

	cellSize := float32(20)

	g := gui.Make(15, 30, gui.MakeOptions{
		CellSize:       &cellSize,
		UpdateCallback: update,
	})

	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
}

func update(matrix gui.Matrix) {
	x := rand.IntN(len(matrix))
	y := rand.IntN(len(matrix[0]))
	matrix[x][y] = rand.IntN(2) > 0
}
