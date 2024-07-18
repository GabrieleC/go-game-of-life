package ebidenUi

import (
	"image/color"

	"gcoletta.it/game-of-life/internal/game"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ebidenUi struct {
	windowTitle         string
	width, height       int
	rows, cols          int
	cellSize, lineWidth float32
	matrix              game.Matrix
	callbacks           game.Callbacks
	cbMap               map[ebiten.Key]func()
	stopRequested       bool
}

var (
	backgroundColor    = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	gridStructureColor = color.RGBA{200, 200, 200, 255}
	cellColor          = color.RGBA{0, 0, 0, 255}
)

func (g *ebidenUi) SetCallbacks(callbacks game.Callbacks) {
	g.callbacks = callbacks
	g.cbMap = map[ebiten.Key]func(){
		ebiten.KeyQ: callbacks.Quit,
		ebiten.KeyL: callbacks.SpeedUp,
		ebiten.KeyK: callbacks.SpeedDown,
		ebiten.KeySpace: callbacks.TogglePlayPause,
	}
}

func (g *ebidenUi) Start() error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.windowTitle)
	return ebiten.RunGame(g)
}

func (g *ebidenUi) Stop() {
	g.stopRequested = true
}

func (g *ebidenUi) UpdateMatrix(update game.MatrixUpdater) {
	n := update(g.matrix)
	if len(n) != len(g.matrix) || (len(n) > 0 && len(n[0]) != len(g.matrix[0])) {
		panic("Matrix dimensions cannot change")
	}
	g.matrix = n
}

func (g ebidenUi) Update() error {
	if g.stopRequested {
		return ebiten.Termination
	}

	for key := range g.cbMap {
		if inpututil.IsKeyJustReleased(key) && g.cbMap[key] != nil {
			g.cbMap[key]()
		}
	}

	return nil
}

func (g ebidenUi) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g ebidenUi) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.drawGridStructure(screen)
	g.drawAliveCells(screen)
}

func (g ebidenUi) drawAliveCells(screen *ebiten.Image) {
	for row, cols := range g.matrix {
		for col, isAlive := range cols {
			if isAlive {
				g.drawAliveCell(screen, row, col)
			}
		}
	}
}

func (g ebidenUi) drawAliveCell(screen *ebiten.Image, row, col int) {
	x := float32(col) * g.cellSize
	y := float32(row) * g.cellSize
	vector.DrawFilledRect(screen, x, y, g.cellSize, g.cellSize, cellColor, true)
}

func (g ebidenUi) drawGridStructure(screen *ebiten.Image) {
	g.drawHorizontalLines(screen)
	g.drawVerticalLines(screen)
}

func (g ebidenUi) drawHorizontalLines(screen *ebiten.Image) {
	bottomY := float32(g.rows) * g.cellSize
	x := float32(0)
	for i := 0; i <= g.cols; i++ {
		vector.StrokeLine(screen, x, 0, x, bottomY, g.lineWidth, gridStructureColor, true)
		x += g.cellSize
	}
}

func (g ebidenUi) drawVerticalLines(screen *ebiten.Image) {
	bottomX := float32(g.cols) * g.cellSize
	y := float32(0)
	for i := 0; i <= g.rows; i++ {
		vector.StrokeLine(screen, 0, y, bottomX, y, g.lineWidth, gridStructureColor, true)
		y += g.cellSize
	}
}
