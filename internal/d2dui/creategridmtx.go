package d2dui

import (
	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/geometry"
)

func (ui *D2dui) makeGridMatrix() grid.Matrix {
	mtx := grid.NewMatrix(ui.matrix.Rows(), ui.matrix.Cols())
	applyEditorPattern(ui.grd, mtx, ui.cursor, ui.editor.currentPattern())
	setAliveCells(ui.matrix, mtx)
	return mtx
}

func applyEditorPattern(grd grid.Grid, mtx grid.Matrix, cursor geometry.Point, pattern game.Matrix) {
	origin, ok := grd.CanvasCoords(cursor)
	if ok {
		grid.ApplyPattern(pattern, origin, mtx, grid.Shadow)
	}
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
