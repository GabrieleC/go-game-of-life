package d2dui

import (
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/patterns"
)

var editorPatterns = [...]game.Matrix{
	nil, //dot
	patterns.Block(),
	patterns.Glider(),
	patterns.LWSS(),
	patterns.MWSS(),
	patterns.HWSS(),
	patterns.Pulsar(),
	patterns.Gosper(),
}

type editor struct {
	editorPatternIdx int
	callback         game.UICallback
}

func (e *editor) iteratePattern() {
	e.editorPatternIdx++
	if e.editorPatternIdx >= len(editorPatterns) {
		e.editorPatternIdx = 0
	}
}

func (e *editor) applyPattern(row, col int) {
	e.callback.Edit(func(matrix game.Matrix) game.Matrix {
		pattern := editorPatterns[e.editorPatternIdx]
		if pattern == nil {
			matrix[row][col] = !matrix[row][col]
		}
		matrix.Copy(pattern, row, col)
		return matrix
	})
}

func (e *editor) currentPattern() game.Matrix {
	return editorPatterns[e.editorPatternIdx]
}
