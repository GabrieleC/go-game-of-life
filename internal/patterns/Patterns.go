package patterns

import (
	_ "embed"

	"gcoletta.it/game-of-life/internal/matrix"
)

//go:embed txt/glider.txt
var glider string

//go:embed txt/pulsar.txt
var pulsar string

func writePattern(dest matrix.Matrix, patternString string, row, col int) {
	pattern := unmarshal(patternString)
	dest.Copy(pattern, row, col)
}

func Glider(matrix matrix.Matrix, row, col int) {
	writePattern(matrix, glider, row, col)
}

func Pulsar(dest matrix.Matrix, row, col int) {
	writePattern(dest, pulsar, row, col)
}
