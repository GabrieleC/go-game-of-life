package grid

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/d2dui/area"
	"gcoletta.it/game-of-life/internal/d2dui/matrixwin"
	"gcoletta.it/game-of-life/internal/game"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
)

type Grid interface {
	Draw(gc *draw2dgl.GraphicContext, origin Coord, canvas area.Area)
	UpdateMatrix(matrix game.Matrix)
	ZoomIn()
	ZoomOut()
	HorizontalPan(steps int)
	VerticalPan(steps int)
}

type GridImpl struct {
	matrix    game.Matrix
	matrixwin *matrixwin.Matrixwin
}

type Coord struct {
	X, Y int
}

func New() Grid {
	return &GridImpl{
		matrixwin: &matrixwin.Matrixwin{},
	}
}

var gridColor = color.RGBA{0, 0, 0, 0x55}

func (g *GridImpl) UpdateMatrix(matrix game.Matrix) {
	if g.matrix == nil {
		g.matrixwin.Update(matrix.Rows(), matrix.Cols())
	}
	g.matrix = matrix
}

func (g *GridImpl) Draw(gc *draw2dgl.GraphicContext, origin Coord, canvas area.Area) {
	if g.matrix != nil {
		width, height := g.calculateGridDimensions(canvas.Width, canvas.Height)
		gridCanvas := area.Area{Width: width, Height: height}

		g.drawGrid(gc, origin, gridCanvas)
		g.drawAliveCells(gc, origin, gridCanvas)
	}
}

func (g *GridImpl) drawGrid(gc *draw2dgl.GraphicContext, origin Coord, canvas area.Area) {
	rows, cols := g.matrixwin.Dimensions()

	cellWidth := canvas.Width / cols
	cellHeight := canvas.Height / rows

	gc.SetStrokeColor(gridColor)

	x := origin.X
	for count := 0; count <= rows; count++ {
		gc.MoveTo(float64(x), float64(origin.Y))
		gc.LineTo(float64(x), float64(origin.Y+(rows*cellHeight)))
		gc.Stroke()
		x += cellWidth
	}

	y := origin.Y
	for count := 0; count <= cols; count++ {
		gc.MoveTo(float64(origin.X), float64(y))
		gc.LineTo(float64(origin.X+(cols*cellWidth)), float64(y))
		gc.Stroke()
		y += cellHeight
	}
}

func (g *GridImpl) drawAliveCells(gc *draw2dgl.GraphicContext, origin Coord, canvas area.Area) {
	maxRow, maxCol := g.matrixwin.Dimensions()

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			matrixRow, matrixCol := g.matrixwin.Coords(row, col)
			if g.matrix[matrixRow][matrixCol] {
				g.DrawAliveCell(gc, origin, canvas, row, col)
			}
		}
	}
}

func (g *GridImpl) DrawAliveCell(gc *draw2dgl.GraphicContext, origin Coord, canvas area.Area, row, col int) {
	rows, cols := g.matrixwin.Dimensions()

	cellWidth := canvas.Width / cols
	cellHeight := canvas.Height / rows

	x := origin.X + cellWidth*col
	y := origin.Y + cellHeight*row

	gc.BeginPath()
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+cellWidth), float64(y+cellHeight))
	gc.SetFillColor(color.RGBA{0, 0, 0, 0xff})
	gc.Fill()
}

func (g *GridImpl) calculateGridDimensions(canvasWidth, canvasHeight int) (int, int) {
	cols, rows := g.matrixwin.Dimensions()
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

func (g *GridImpl) ZoomIn() {
	g.matrixwin.ZoomIn()
}

func (g *GridImpl) ZoomOut() {
	g.matrixwin.ZoomOut()
}

func (g *GridImpl) HorizontalPan(steps int) {
	g.matrixwin.HorizontalPan(steps)
}

func (g *GridImpl) VerticalPan(steps int) {
	g.matrixwin.VerticalPan(steps)
}
