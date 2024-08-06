package grid

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/d2dui/area"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
)

var gridColor = color.RGBA{0, 0, 0, 0x55}
var aliveCellColor = color.RGBA{0, 0, 0, 0xff}
var shadowCellColor = color.RGBA{0, 0, 0, 0x33}

func Draw(gc *draw2dgl.GraphicContext, matrix Matrix, canvas area.Area) {
	if matrix != nil {
		width, height := calculateGridDimensions(matrix, canvas.Width, canvas.Height)
		gridCanvas := area.Area{Width: width, Height: height}

		drawGrid(gc, matrix, gridCanvas)
		drawAliveCells(gc, matrix, gridCanvas)
	}
}

func drawGrid(gc *draw2dgl.GraphicContext, matrix Matrix, canvas area.Area) {
	rows, cols := dimension(matrix)

	cellWidth := canvas.Width / cols
	cellHeight := canvas.Height / rows

	gc.SetStrokeColor(gridColor)

	for x, count := 0, 0; count <= rows; count++ {
		gc.MoveTo(float64(x), 0)
		gc.LineTo(float64(x), float64(rows*cellHeight))
		gc.Stroke()
		x += cellWidth
	}

	for y, count := 0, 0; count <= cols; count++ {
		gc.MoveTo(0, float64(y))
		gc.LineTo(float64(cols*cellWidth), float64(y))
		gc.Stroke()
		y += cellHeight
	}
}

func drawAliveCells(gc *draw2dgl.GraphicContext, matrix Matrix, canvas area.Area) {
	maxRow, maxCol := dimension(matrix)

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			matrixRow, matrixCol := row, col
			switch matrix[matrixRow][matrixCol] {
			case Alive:
				drawCell(gc, matrix, canvas, row, col, aliveCellColor)
			case Shadow:
				drawCell(gc, matrix, canvas, row, col, shadowCellColor)
			}
		}
	}
}

func drawCell(gc *draw2dgl.GraphicContext, matrix Matrix, canvas area.Area, row, col int, color color.Color) {
	rows, cols := dimension(matrix)

	cellWidth := canvas.Width / cols
	cellHeight := canvas.Height / rows

	x := cellWidth * col
	y := cellHeight * row

	gc.BeginPath()
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+cellWidth), float64(y+cellHeight))
	gc.SetFillColor(color)
	gc.Fill()
}

func calculateGridDimensions(matrix Matrix, canvasWidth, canvasHeight int) (int, int) {
	rows, cols := dimension(matrix)
	var width, height int

	tallerThanWiderComparision := area.IsTallerThanWider(
		area.Area{Width: canvasWidth, Height: canvasHeight},
		area.Area{Width: cols, Height: rows},
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
