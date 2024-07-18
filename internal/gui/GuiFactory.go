package gui

import (
	"gcoletta.it/game-of-life/internal/matrix"
	"gcoletta.it/game-of-life/internal/utils"
)

type Options struct {
	WindowTitle *string
	CellSize    float32
	LineWidth   float32
}

func New(rows, cols int, opts Options) Gui {

	const defaultCellSize = 20
	const defaultLineWidth = 0.5
	const defaultWindowTitle = ""

	cellSize := utils.CoalesceF32(opts.CellSize, defaultCellSize)

	g := guiImpl{
		width:       cols * int(cellSize),
		height:      rows * int(cellSize),
		rows:        rows,
		cols:        cols,
		windowTitle: utils.CoalescePtr(opts.WindowTitle, defaultWindowTitle),
		cellSize:    utils.CoalesceF32(opts.CellSize, defaultCellSize),
		lineWidth:   utils.CoalesceF32(opts.LineWidth, defaultLineWidth),
		matrix:      Matrix(matrix.NewMatrix(rows, cols)),
	}

	return &g
}
