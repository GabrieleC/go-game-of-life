package main

import (
	"log"
	"math/rand/v2"

	"gcoletta.it/game-of-life/gui"
)

func main() {

	g := gui.Make(15, 30, gui.MakeOptions{
		CellSize:       20,
		UpdateCallback: onGuiUpdate,
	})

	startGui(g)
}

func startGui(g gui.Gui) {
	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
}

func onGuiUpdate(matrix gui.Matrix) {
	x := rand.IntN(len(matrix))
	y := rand.IntN(len(matrix[0]))
	matrix[x][y] = rand.IntN(2) > 0
}
