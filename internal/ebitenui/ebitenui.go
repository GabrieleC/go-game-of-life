package ebitenui

import (
	"image/color"
	"time"

	"gcoletta.it/game-of-life/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type keyCallbackMap map[ebiten.Key]func()

type ebitenUi struct {
	windowTitle         string
	width, height       int
	rows, cols          int
	cellSize, lineWidth float32
	matrix              game.Matrix
	oneShotKeyMap       keyCallbackMap
	longPressKeyMap     keyCallbackMap
	stopRequested       bool
	keyPressTimestamps  map[ebiten.Key]int64
}

var (
	backgroundColor    = color.NRGBA{0xff, 0xff, 0xff, 0xff}
	gridStructureColor = color.RGBA{200, 200, 200, 255}
	cellColor          = color.RGBA{0, 0, 0, 255}
)

func (g *ebitenUi) SetGame(callback game.Game) {
	g.oneShotKeyMap = map[ebiten.Key]func(){
		ebiten.KeyQ:     callback.Quit,
		ebiten.KeySpace: callback.TogglePlayPause,
	}
	g.longPressKeyMap = map[ebiten.Key]func(){
		ebiten.KeyUp:    callback.SpeedUp,
		ebiten.KeyDown:  callback.SpeedDown,
		ebiten.KeyLeft:  callback.Back,
		ebiten.KeyRight: callback.Next,
	}
}

func (g *ebitenUi) Start() error {
	ebiten.SetWindowSize(g.width, g.height)
	ebiten.SetWindowTitle(g.windowTitle)
	return ebiten.RunGame(g)
}

func (g *ebitenUi) Stop() {
	g.stopRequested = true
}

func (g *ebitenUi) UpdateMatrix(update game.MatrixUpdater) {
	n := update(g.matrix)
	if len(n) != len(g.matrix) || (len(n) > 0 && len(n[0]) != len(g.matrix[0])) {
		panic("Matrix dimensions cannot change")
	}
	g.matrix = n
}

func (g ebitenUi) Update() error {
	if g.stopRequested {
		return ebiten.Termination
	}

	for key := range g.oneShotKeyMap {
		if inpututil.IsKeyJustReleased(key) && g.oneShotKeyMap[key] != nil {
			g.oneShotKeyMap[key]()
		}
	}

	for key := range g.longPressKeyMap {
		if g.isKeyPressed(key) && g.longPressKeyMap[key] != nil {
			g.longPressKeyMap[key]()
		}
	}

	return nil
}

func (g ebitenUi) isKeyPressed(key ebiten.Key) bool {
	now := time.Now().UnixMilli()
	if ebiten.IsKeyPressed(key) && now-g.keyPressTimestamps[key] > 70 {
		g.keyPressTimestamps[key] = now
		return true
	}
	return false
}

func (g ebitenUi) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width, g.height
}

func (g ebitenUi) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	g.drawGridStructure(screen)
	g.drawAliveCells(screen)
}

func (g ebitenUi) drawAliveCells(screen *ebiten.Image) {
	for row, cols := range g.matrix {
		for col, isAlive := range cols {
			if isAlive {
				g.drawAliveCell(screen, row, col)
			}
		}
	}
}

func (g ebitenUi) drawAliveCell(screen *ebiten.Image, row, col int) {
	x := float32(col) * g.cellSize
	y := float32(row) * g.cellSize
	vector.DrawFilledRect(screen, x, y, g.cellSize, g.cellSize, cellColor, true)
}

func (g ebitenUi) drawGridStructure(screen *ebiten.Image) {
	g.drawHorizontalLines(screen)
	g.drawVerticalLines(screen)
}

func (g ebitenUi) drawHorizontalLines(screen *ebiten.Image) {
	bottomY := float32(g.rows) * g.cellSize
	x := float32(0)
	for i := 0; i <= g.cols; i++ {
		vector.StrokeLine(screen, x, 0, x, bottomY, g.lineWidth, gridStructureColor, true)
		x += g.cellSize
	}
}

func (g ebitenUi) drawVerticalLines(screen *ebiten.Image) {
	bottomX := float32(g.cols) * g.cellSize
	y := float32(0)
	for i := 0; i <= g.rows; i++ {
		vector.StrokeLine(screen, 0, y, bottomX, y, g.lineWidth, gridStructureColor, true)
		y += g.cellSize
	}
}
