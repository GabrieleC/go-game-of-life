package grid

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/geometry"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
)

var gridColor = color.RGBA{0, 0, 0, 0x55}
var alivecolor = color.RGBA{0, 0, 0, 0xff}
var shadowcolor = color.RGBA{0, 0, 0, 0x33}

func stateColor(state byte) color.Color {
	switch state {
	case Alive:
		return alivecolor
	case Shadow:
		return shadowcolor
	default:
		return nil
	}
}

func CanvasCoords(x, y float64, canvas geometry.Area, maxRow, maxCol int, origin geometry.Point, cellSize int) (row, col int, ok bool) {
	row = int((y + float64(origin.Y)) / float64(cellSize))
	col = int((x + float64(origin.X)) / float64(cellSize))

	if row <= maxRow && col <= maxCol {
		return row, col, true
	} else {
		return 0, 0, false
	}
}

func Draw(gc *draw2dgl.GraphicContext, matrix Matrix, canvas geometry.Area, origin geometry.Point, cellSize int) {
	if matrix != nil {
		drawGrid(gc, matrix, canvas, origin, cellSize)
		drawAliveCells(gc, matrix, origin, cellSize)
	}
}

func drawGrid(gc *draw2dgl.GraphicContext, matrix Matrix, canvas geometry.Area, origin geometry.Point, cellSize int) {
	rows, cols := dimension(matrix)
	gc.SetStrokeColor(gridColor)

	firstCol := origin.X / cellSize
	firstRow := origin.Y / cellSize

	gridHeight := (rows - firstRow) * cellSize
	xCursor := cellSize - (origin.X % cellSize)
	colCur := firstCol
	for xCursor <= canvas.Width && colCur <= cols - 1 {
		gc.MoveTo(float64(xCursor), 0)
		gc.LineTo(float64(xCursor), float64(gridHeight))
		gc.Stroke()
		colCur++
		xCursor += cellSize
	}

	gridWidth := (cols - firstCol) * cellSize
	yCursor := cellSize - (origin.Y % cellSize)
	rowcur := firstRow
	for yCursor <= canvas.Height && rowcur <= rows - 1 {
		gc.MoveTo(0, float64(yCursor))
		gc.LineTo(float64(gridWidth), float64(yCursor))
		gc.Stroke()
		rowcur++
		yCursor += cellSize
	}
}

func drawAliveCells(gc *draw2dgl.GraphicContext, matrix Matrix, origin geometry.Point, cellSize int) {
	maxRow, maxCol := dimension(matrix)

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {

			cellOrigin := geometry.Point{
				X: (col * cellSize) - origin.X,
				Y: (row * cellSize) - origin.Y,
			}

			if matrix[row][col] != Dead {
				color := stateColor(matrix[row][col])
				drawCell(gc, color, cellOrigin, cellSize)
			}
		}
	}
}

func drawCell(gc *draw2dgl.GraphicContext, color color.Color, cellorigin geometry.Point, cellsize int) {
	gc.BeginPath()
	draw2dkit.Rectangle(gc,
		float64(cellorigin.X),
		float64(cellorigin.Y),
		float64(cellorigin.X+cellsize),
		float64(cellorigin.Y+cellsize))
	gc.SetFillColor(color)
	gc.Fill()
}
