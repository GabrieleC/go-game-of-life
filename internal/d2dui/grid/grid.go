package grid

import (
	"image/color"

	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
)

type Description struct {
	OriginX, OriginY int
	Width, Height    int
	Cols, Rows       int
}

var gridColor = color.RGBA{0, 0, 0, 0x55}

func Draw(gc *draw2dgl.GraphicContext, g Description) {

	cellWidth := g.Width / g.Cols
	cellHeight := g.Height / g.Rows

	gc.SetStrokeColor(gridColor)

	x := g.OriginX
	for count := 0; count <= g.Rows; count++ {
		gc.MoveTo(float64(x), float64(g.OriginY))
		gc.LineTo(float64(x), float64(g.OriginY+(g.Rows*cellHeight)))
		gc.Stroke()
		x += cellWidth
	}

	y := g.OriginY
	for count := 0; count <= g.Cols; count++ {
		gc.MoveTo(float64(g.OriginX), float64(y))
		gc.LineTo(float64(g.OriginX+(g.Cols*cellWidth)), float64(y))
		gc.Stroke()
		y += cellHeight
	}
}

func DrawAliveCell(gc *draw2dgl.GraphicContext, g Description, row, col int) {

	cellWidth := g.Width / g.Cols
	cellHeight := g.Height / g.Rows

	x := g.OriginX + cellWidth*col
	y := g.OriginY + cellHeight*row

	gc.BeginPath()
	draw2dkit.Rectangle(gc, float64(x), float64(y), float64(x+cellWidth), float64(y+cellHeight))
	gc.SetFillColor(color.RGBA{0, 0, 0, 0xff})
	gc.Fill()
}
