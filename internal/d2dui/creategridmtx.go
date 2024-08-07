package d2dui

import (
	"gcoletta.it/game-of-life/internal/d2dui/area"
	"gcoletta.it/game-of-life/internal/d2dui/grid"
)

func (ui *D2dui) createGridMatrix() grid.Matrix {
	rows, cols := ui.matrixwin.Dimensions()
	mtx := newGridMatrix(rows, cols)
	ui.setShadowCells(rows, cols, mtx)
	ui.setAliveCells(mtx)
	return mtx
}

func newGridMatrix(rows int, cols int) [][]byte {
	mtx := make([][]byte, rows)
	for rowId := range mtx {
		mtx[rowId] = make([]byte, cols)
	}
	return mtx
}

func (ui *D2dui) setShadowCells(rows int, cols int, mtx [][]byte) {
	gridArea := area.Area{Width: ui.width, Height: ui.height}
	curRow, curCol, ok := grid.CanvasCoords(ui.curX, ui.curY, gridArea, rows, cols)
	if ok {
		mtx[curRow][curCol] = grid.Shadow
	}
}

func (ui *D2dui) setAliveCells(mtx [][]byte) {
	originRow, originCol := ui.matrixwin.Origin()
	for rowId, row := range mtx {
		for colId := range row {
			if ui.matrix[rowId+originRow][colId+originCol] {
				mtx[rowId][colId] = grid.Alive
			}
		}
	}
}
