package gui

import "gcoletta.it/game-of-life/utils"

type MakeOptions struct {
	WindowTitle    *string
	CellSize       float32
	LineWidth      float32
	UpdateCallback updateCallback
}

func Make(rows, cols int, opts MakeOptions) Gui {

	const defaultCellSize = 20
	const defaultLineWidth = 0.5
	const defaultWindowTitle = ""

	cellSize := utils.CoalesceF32(opts.CellSize, defaultCellSize)

	g := Gui{
		width:          cols * int(cellSize),
		height:         rows * int(cellSize),
		rows:           rows,
		cols:           cols,
		windowTitle:    utils.CoalescePtr(opts.WindowTitle, defaultWindowTitle),
		cellSize:       utils.CoalesceF32(opts.CellSize, defaultCellSize),
		lineWidth:      utils.CoalesceF32(opts.LineWidth, defaultLineWidth),
		matrix:         utils.InitMatrix(rows, cols),
		updateCallback: opts.UpdateCallback,
	}

	return g
}
