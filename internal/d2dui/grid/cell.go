package grid

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/geometry"
	"github.com/llgcode/draw2d/draw2dgl"
	"github.com/llgcode/draw2d/draw2dkit"
)

type cell struct {
	State  byte
	Origin geometry.Point
	Size   int
}

const (
	Dead = iota
	Alive
	Shadow
)

var (
	aliveColor  = color.RGBA{0, 0, 0, 0xff}
	shadowColor = color.RGBA{0, 0, 0, 0x33}
)

func colorByType(state byte) color.Color {
	switch state {
	case Alive:
		return aliveColor
	case Shadow:
		return shadowColor
	default:
		return nil
	}
}

func drawCell(gc *draw2dgl.GraphicContext, cell cell) {
	color := colorByType(cell.State)
	gc.BeginPath()
	draw2dkit.Rectangle(gc,
		float64(cell.Origin.X),
		float64(cell.Origin.Y),
		float64(cell.Origin.X+cell.Size),
		float64(cell.Origin.Y+cell.Size))
	gc.SetFillColor(color)
	gc.Fill()
}
