package gui

type MakeOptions struct {
	WindowTitle    *string
	CellSize       *float32
	LineWidth      *float32
	UpdateCallback updateCallback
}

func Make(rows, cols int, opts MakeOptions) Gui {

	const defaultCellSize = 20
	const defaultLineWidth = 0.5
	const defaultWindowTitle = ""

	cellSize := coalesce(opts.CellSize, defaultCellSize)

	g := Gui{
		width:          cols * int(cellSize),
		height:         rows * int(cellSize),
		rows:           rows,
		cols:           cols,
		windowTitle:    coalesce(opts.WindowTitle, defaultWindowTitle),
		cellSize:       coalesce(opts.CellSize, defaultCellSize),
		lineWidth:      coalesce(opts.LineWidth, defaultLineWidth),
		matrix:         initMatrix(rows, cols),
		updateCallback: opts.UpdateCallback,
	}

	return g
}
