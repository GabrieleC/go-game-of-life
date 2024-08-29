package grid

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/geometry"
	"github.com/llgcode/draw2d/draw2dgl"
)

var gridColor = color.RGBA{0, 0, 0, 0x55}

const panAmount = 5
const cellMaxSize = 100

type Grid struct {
	Matrix   Matrix
	Canvas   geometry.Area
	Origin   geometry.Point
	CellSize int
}

func (grid Grid) CanvasCoords(point geometry.Point) (geometry.Point, bool) {
	row := (point.Y + grid.Origin.Y) / grid.CellSize
	col := (point.X + grid.Origin.X) / grid.CellSize

	maxRows, maxCols := Dimension(grid.Matrix)
	if row > maxRows || col > maxCols {
		return geometry.Point{}, false
	}

	return geometry.Point{X: col, Y: row}, true
}

func (grid Grid) Draw(gc *draw2dgl.GraphicContext) {
	if grid.Matrix != nil {
		drawGrid(gc, grid)
		drawCells(gc, grid)
	}
}

func (grd *Grid) ZoomOut() {
	grd.CellSize = max(grd.CellSize-1, 3)
}

func (grd *Grid) ZoomIn() {
	grd.CellSize = min(grd.CellSize+1, cellMaxSize)
}

func (grd *Grid) PanUp() {
	grd.Origin.Y = max(grd.Origin.Y-grd.CellSize, 0)
}

func (grd *Grid) PanDown() {
	rows, _ := Dimension(grd.Matrix)
	maxy := (grd.CellSize * rows) - grd.Canvas.Height
	grd.Origin.Y = min(grd.Origin.Y+grd.CellSize, maxy)
}

func (grd *Grid) PanLeft() {
	grd.Origin.X = max(grd.Origin.X-grd.CellSize, 0)
}

func (grd *Grid) PanRight() {
	_, cols := Dimension(grd.Matrix)
	maxx := (grd.CellSize * cols) - grd.Canvas.Width
	grd.Origin.X = min(grd.Origin.X+grd.CellSize, maxx)
}

func drawGrid(gc *draw2dgl.GraphicContext, grid Grid) {
	rows, cols := Dimension(grid.Matrix)
	gc.SetStrokeColor(gridColor)

	firstCol := grid.Origin.X / grid.CellSize
	firstRow := grid.Origin.Y / grid.CellSize

	gridHeight := (rows - firstRow) * grid.CellSize
	xCursor := grid.CellSize - (grid.Origin.X % grid.CellSize)
	colCur := firstCol
	for xCursor <= grid.Canvas.Width && colCur <= cols-1 {
		gc.MoveTo(float64(xCursor), 0)
		gc.LineTo(float64(xCursor), float64(gridHeight))
		gc.Stroke()
		colCur++
		xCursor += grid.CellSize
	}

	gridWidth := (cols - firstCol) * grid.CellSize
	yCursor := grid.CellSize - (grid.Origin.Y % grid.CellSize)
	rowcur := firstRow
	for yCursor <= grid.Canvas.Height && rowcur <= rows-1 {
		gc.MoveTo(0, float64(yCursor))
		gc.LineTo(float64(gridWidth), float64(yCursor))
		gc.Stroke()
		rowcur++
		yCursor += grid.CellSize
	}
}

func drawCells(gc *draw2dgl.GraphicContext, grid Grid) {
	maxRow, maxCol := Dimension(grid.Matrix)

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {

			if grid.Matrix[row][col] == Dead {
				continue
			}

			cellOrigin := geometry.Point{
				X: (col * grid.CellSize) - grid.Origin.X,
				Y: (row * grid.CellSize) - grid.Origin.Y,
			}

			cell := cell{
				State:  grid.Matrix[row][col],
				Origin: cellOrigin,
				Size:   grid.CellSize,
			}

			drawCell(gc, cell)
		}
	}
}
