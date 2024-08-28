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

	colcur := origin.X / cellSize
	rowcur := origin.Y / cellSize

	xcursor := cellSize - (origin.X % cellSize)
	for xcursor <= canvas.Width && colcur <= cols {
		gc.MoveTo(float64(xcursor), 0)
		gc.LineTo(float64(xcursor), float64((rows-rowcur)*cellSize))
		gc.Stroke()
		colcur++
		xcursor += cellSize
	}

	ycursor := cellSize - (origin.Y % cellSize)
	for ycursor <= canvas.Height && rowcur <= rows {
		gc.MoveTo(0, float64(ycursor))
		gc.LineTo(float64((cols-colcur)*cellSize), float64(ycursor))
		gc.Stroke()
		rowcur++
		ycursor += cellSize
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

func calculateGridDimensions(rows, cols int, canvasWidth, canvasHeight int) (int, int) {
	var width, height int

	tallerThanWiderComparision := geometry.IsTallerThanWider(
		geometry.Area{Width: canvasWidth, Height: canvasHeight},
		geometry.Area{Width: cols, Height: rows},
	)

	if tallerThanWiderComparision > 0 {
		matrixRatio := float64(cols) / float64(rows)
		width = int(float64(canvasHeight) * matrixRatio)
		height = canvasHeight
	} else {
		matrixRatio := float64(rows) / float64(cols)
		height = int(float64(canvasWidth) * matrixRatio)
		width = canvasWidth
	}

	return width, height
}
