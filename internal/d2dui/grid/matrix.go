package grid

import (
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/geom"
)

type Matrix [][]byte

func NewMatrix(rows int, cols int) Matrix {
	mtx := make(Matrix, rows)
	for rowId := range mtx {
		mtx[rowId] = make([]byte, cols)
	}
	return mtx
}

func Dimension(matrix Matrix) (rows, cols int) {
	rows = len(matrix)
	if rows > 0 {
		cols = len(matrix[0])
	}
	return rows, cols
}

type StateUpdater func(currentState byte) byte

func ApplyPattern(pattern game.Matrix, origin geom.Point, mtx Matrix, updater StateUpdater) {
	rows, cols := Dimension(mtx)
	for row := 0; row < pattern.Rows() && row+origin.Y < rows; row++ {
		for col := 0; col < pattern.Cols() && col+origin.X < cols; col++ {
			if pattern[row][col] {
				mtx[row+origin.Y][col+origin.X] = updater(mtx[row+origin.Y][col+origin.X])
			}
		}
	}
}
