package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type guiImpl struct {
	windowTitle         string
	width, height       int
	rows, cols          int
	cellSize, lineWidth float32
	matrix              Matrix
	callbacks           Callbacks
	cbMap				map[ebiten.Key]func()
	stopRequested       bool
}

var (
	backgroundColor    = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	gridStructureColor = color.RGBA{200, 200, 200, 255}
	cellColor          = color.RGBA{0, 0, 0, 255}
)

func (g *guiImpl) SetCallbacks(callbacks Callbacks) {
	g.callbacks = callbacks
	g.cbMap = map[ebiten.Key]func(){
		ebiten.KeyQ: callbacks.Quit,
		ebiten.KeyL: callbacks.SpeedUp,
		ebiten.KeyK: callbacks.SpeedDown,
	}
}

func (g *guiImpl) Start() error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.windowTitle)
	return ebiten.RunGame(g)
}

func (g *guiImpl) Stop() {
	g.stopRequested = true
}

func (g *guiImpl) UpdateMatrix(update MatrixUpdater) {
	n := update(g.matrix)
	if len(n) != len(g.matrix) || (len(n) > 0 && len(n[0]) != len(g.matrix[0])) {
		panic("Matrix dimensions cannot change")
	}
	g.matrix = n
}

func (g guiImpl) Update() error {
	if g.stopRequested {
		return ebiten.Termination
	}

	for key := range g.cbMap {
		if ebiten.IsKeyPressed(key) && g.cbMap[key] != nil {
			g.cbMap[key]()
		}
	}

	return nil
}

func (g guiImpl) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g guiImpl) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.drawGridStructure(screen)
	g.drawAliveCells(screen)
}

func (g guiImpl) drawAliveCells(screen *ebiten.Image) {
	for row, cols := range g.matrix {
		for col, isAlive := range cols {
			if isAlive {
				g.drawAliveCell(screen, row, col)
			}
		}
	}
}

func (g guiImpl) drawAliveCell(screen *ebiten.Image, row, col int) {
	x := float32(col) * g.cellSize
	y := float32(row) * g.cellSize
	vector.DrawFilledRect(screen, x, y, g.cellSize, g.cellSize, cellColor, true)
}

func (g guiImpl) drawGridStructure(screen *ebiten.Image) {
	g.drawHorizontalLines(screen)
	g.drawVerticalLines(screen)
}

func (g guiImpl) drawHorizontalLines(screen *ebiten.Image) {
	bottomY := float32(g.rows) * g.cellSize
	x := float32(0)
	for i := 0; i <= g.cols; i++ {
		vector.StrokeLine(screen, x, 0, x, bottomY, g.lineWidth, gridStructureColor, true)
		x += g.cellSize
	}
}

func (g guiImpl) drawVerticalLines(screen *ebiten.Image) {
	bottomX := float32(g.cols) * g.cellSize
	y := float32(0)
	for i := 0; i <= g.rows; i++ {
		vector.StrokeLine(screen, 0, y, bottomX, y, g.lineWidth, gridStructureColor, true)
		y += g.cellSize
	}
}
