package grid

import (
	"gcoletta.it/game-of-life/internal/geom"
	"gcoletta.it/game-of-life/internal/matrix"
	"gcoletta.it/game-of-life/internal/patterns"
)

type GridMatrix matrix.Matrix[CellState]

func NewGridMatrix(rows int, cols int) GridMatrix {
	return GridMatrix(matrix.Create[CellState](rows, cols))
}

func Dimensions(mtx GridMatrix) (rows, cols int) {
	return matrix.Dimensions(matrix.Matrix[CellState](mtx))
}

func Copy(mtx GridMatrix) GridMatrix {
	copied := matrix.Copy(matrix.Matrix[CellState](mtx))
	return GridMatrix(copied)
}

type StateUpdater func(currentState CellState) CellState

func ApplyPattern(pattern patterns.Pattern, origin geom.Point, mtx GridMatrix, updater StateUpdater) {
	rows, cols := matrix.Dimensions(matrix.Matrix[CellState](mtx))
	patternRows, patternCols := matrix.Dimensions(matrix.Matrix[bool](pattern))
	for row := 0; row < patternRows && row+origin.Y < rows; row++ {
		for col := 0; col < patternCols && col+origin.X < cols; col++ {
			if pattern[row][col] {
				mtx[row+origin.Y][col+origin.X] = updater(mtx[row+origin.Y][col+origin.X])
			}
		}
	}
}
