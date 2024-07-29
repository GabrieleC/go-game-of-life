package ebitenui

import (
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Options struct {
	WindowTitle *string
	CellSize    float32
	LineWidth   float32
}

func New(rows, cols int, opts Options) game.Gui {

	const defaultCellSize = 20
	const defaultLineWidth = 0.5
	const defaultWindowTitle = ""

	cellSize := utils.CoalesceF32(opts.CellSize, defaultCellSize)

	g := ebitenUi{
		width:              cols * int(cellSize),
		height:             rows * int(cellSize),
		rows:               rows,
		cols:               cols,
		windowTitle:        utils.CoalescePtr(opts.WindowTitle, defaultWindowTitle),
		cellSize:           utils.CoalesceF32(opts.CellSize, defaultCellSize),
		lineWidth:          utils.CoalesceF32(opts.LineWidth, defaultLineWidth),
		matrix:             game.NewMatrix(rows, cols),
		keyPressTimestamps: map[ebiten.Key]int64{},
	}

	return &g
}
