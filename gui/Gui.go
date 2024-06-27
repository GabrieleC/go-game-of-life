package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Gui struct {
	screenWidth, screenHeight int
	windowTitle               string
	canvas                    canvas
}

type Matrix [][]bool

type MakeOptions struct {
	WindowTitle *string
	CellSize    *float32
	LineWidth   *float32
}

func Make(gridSize int, opts MakeOptions) (Gui, Matrix) {

	const defaultCellSize = 20
	const defaultLineWidth = 0.5
	const defaultWindowTitle = ""

	title := defaultWindowTitle
	if opts.WindowTitle != nil {
		title = *opts.WindowTitle
	}

	cellSize := float32(defaultCellSize)
	if opts.CellSize != nil {
		cellSize = *opts.CellSize
	}
	
	lineWidth := float32(defaultLineWidth)
	if opts.LineWidth != nil {
		lineWidth = *opts.LineWidth
	}

	screenSize := gridSize * int(cellSize)

	g := Gui{screenSize, screenSize, title,
		canvas{
			width:          screenSize,
			height:         screenSize,
			matrix:         initMatrix(gridSize, gridSize),
			cellSize:       cellSize,
			lineWidth:      lineWidth,
			updateCallback: onUpdate,
		}}

	return g, g.canvas.matrix
}

func (g *Gui) Start() error {
	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle(g.windowTitle)
	return ebiten.RunGame(&g.canvas)
}

func initMatrix(rows, cols int) Matrix {

	matrix := make([][]bool, rows)
	for i := range matrix {
		matrix[i] = make([]bool, cols)
	}

	return matrix
}

func onUpdate() error {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	return nil
}
