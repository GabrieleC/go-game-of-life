package adapters

import (
	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
)

func makeGridMatrix(matrix game.Matrix) grid.Matrix {
	mtx := grid.NewMatrix(matrix.Rows(), matrix.Cols())
	setAliveCells(matrix, mtx)
	return mtx
}

func setAliveCells(matrix game.Matrix, mtx grid.Matrix) {
	for rowId, row := range mtx {
		for colId := range row {
			if matrix[rowId][colId] {
				mtx[rowId][colId] = grid.Alive
			}
		}
	}
}
