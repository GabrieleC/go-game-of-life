package d2dui

import (
	"gcoletta.it/game-of-life/internal/d2dui/area"
	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
)

func (ui *D2dui) createGridMatrix() grid.Matrix {
	cols := ui.matrix.Cols()
	rows := ui.matrix.Rows()
	mtx := newGridMatrix(rows, cols)
	ui.drawEditorPattern(rows, cols, mtx)
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

func (ui *D2dui) drawEditorPattern(rows int, cols int, mtx [][]byte) {
	gridArea := area.Area{Width: ui.width, Height: ui.height}
	curRow, curCol, ok := grid.CanvasCoords(ui.curX, ui.curY, gridArea, rows, cols)
	if ok {
		pattern := ui.editor.currentPattern()
		ui.copyPattern(pattern, curRow, curCol, rows, cols, mtx)
	}
}

func (*D2dui) copyPattern(pattern game.Matrix, originRow, originCol int, rows, cols int, mtx [][]byte) {
	for row := 0; row < pattern.Rows() && row+originRow < rows; row++ {
		for col := 0; col < pattern.Cols() && col+originCol < cols; col++ {
			if pattern[row][col] {
				mtx[row+originRow][col+originCol] = grid.Shadow
			}
		}
	}
}

func (ui *D2dui) setAliveCells(mtx [][]byte) {
	for rowId, row := range mtx {
		for colId := range row {
			if ui.matrix[rowId][colId] {
				mtx[rowId][colId] = grid.Alive
			}
		}
	}
}
