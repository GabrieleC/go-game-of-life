package gui

import (
	"fmt"
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var backgroundColor = color.NRGBA{0xff, 0xff, 0xff, 0xff}
var gridStructureColor = color.RGBA{200, 200, 200, 255}
var cellColor = color.RGBA{0, 0, 0, 255}

type updateCallback func() error

type canvas struct {
	width, height                int
	matrix                       [][]bool
	padding, cellSize, lineWidth float32
	updateCallback
}

func (c *canvas) Update() error {
	return c.updateCallback()
}

func (g *canvas) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

var first = 0

func (c *canvas) Draw(screen *ebiten.Image) {
	if first%60 == 0 {
		x := rand.IntN(len(c.matrix))
		y := rand.IntN(len(c.matrix))
		c.matrix[x][y] = rand.IntN(2) > 0
	}
	first++

	screen.Fill(backgroundColor)
	c.drawGridStructure(screen)
	c.drawAliveCells(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (c *canvas) drawAliveCells(screen *ebiten.Image) {
	for row, cols := range c.matrix {
		for col, isAlive := range cols {
			if isAlive {
				c.drawAliveCell(screen, row, col)
			}
		}
	}
}

func (c *canvas) drawAliveCell(screen *ebiten.Image, col, row int) {
	x := (float32(col) * c.cellSize) + c.padding
	y := (float32(row) * c.cellSize) + c.padding
	vector.DrawFilledRect(screen, x, y, c.cellSize, c.cellSize, cellColor, true)
}

func (c *canvas) drawGridStructure(screen *ebiten.Image) {

	rows, cols := matrixSizes(c.matrix)

	bottomY := c.padding + (float32(rows) * c.cellSize)
	x := float32(c.padding)
	for i := 0; i <= cols; i++ {
		vector.StrokeLine(screen, x, c.padding, x, bottomY, c.lineWidth, gridStructureColor, true)
		x += float32(c.cellSize)
	}

	bottomX := c.padding + (float32(cols) * c.cellSize)
	y := float32(c.padding)
	for i := 0; i <= rows; i++ {
		vector.StrokeLine(screen, c.padding, y, bottomX, y, c.lineWidth, gridStructureColor, true)
		y += float32(c.cellSize)
	}
}

func matrixSizes(matrix [][]bool) (rows int, cols int) {
	if matrix != nil {
		rows = len(matrix)
		if rows > 0 {
			cols = len(matrix[0])
		}
	}
	return
}
