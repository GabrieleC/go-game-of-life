package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Gui struct {
	windowTitle         string
	width, height       int
	rows, cols          int
	cellSize, lineWidth float32
	matrix              Matrix
	updateCallback      updateCallback
}

type Matrix [][]bool
type updateCallback func(matrix Matrix)

var (
	backgroundColor    = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	gridStructureColor = color.RGBA{200, 200, 200, 255}
	cellColor          = color.RGBA{0, 0, 0, 255}
)

func (g *Gui) Start() error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.windowTitle)
	return ebiten.RunGame(g)
}

func (g *Gui) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	if g.updateCallback != nil {
		g.updateCallback(g.matrix)
	}
	return nil
}

func (g *Gui) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g *Gui) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.drawGridStructure(screen)
	g.drawAliveCells(screen)
}

func (g *Gui) drawAliveCells(screen *ebiten.Image) {
	for row, cols := range g.matrix {
		for col, isAlive := range cols {
			if isAlive {
				g.drawAliveCell(screen, row, col)
			}
		}
	}
}

func (c *Gui) drawAliveCell(screen *ebiten.Image, row, col int) {
	x := float32(col) * c.cellSize
	y := float32(row) * c.cellSize
	vector.DrawFilledRect(screen, x, y, c.cellSize, c.cellSize, cellColor, true)
}

func (g *Gui) drawGridStructure(screen *ebiten.Image) {
	g.drawHorizontalLines(screen)
	g.drawVerticalLines(screen)
}

func (g *Gui) drawHorizontalLines(screen *ebiten.Image) {
	bottomY := float32(g.rows) * g.cellSize
	x := float32(0)
	for i := 0; i <= g.cols; i++ {
		vector.StrokeLine(screen, x, 0, x, bottomY, g.lineWidth, gridStructureColor, true)
		x += g.cellSize
	}
}

func (g *Gui) drawVerticalLines(screen *ebiten.Image) {
	bottomX := float32(g.cols) * g.cellSize
	y := float32(0)
	for i := 0; i <= g.rows; i++ {
		vector.StrokeLine(screen, 0, y, bottomX, y, g.lineWidth, gridStructureColor, true)
		y += g.cellSize
	}
}
