package patterns

import (
	_ "embed"

	"gcoletta.it/game-of-life/internal/game"
)

//go:embed txt/glider.txt
var glider string

//go:embed txt/pulsar.txt
var pulsar string

//go:embed txt/block.txt
var block string

func writePattern(dest game.Matrix, patternString string, row, col int) {
	pattern := unmarshal(patternString)
	dest.Copy(pattern, row, col)
}

func Glider(matrix game.Matrix, row, col int) {
	writePattern(matrix, glider, row, col)
}

func Pulsar(dest game.Matrix, row, col int) {
	writePattern(dest, pulsar, row, col)
}

func Block(dest game.Matrix, row, col int) {
	writePattern(dest, block, row, col)
}
