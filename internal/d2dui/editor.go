package d2dui

import (
	"gcoletta.it/game-of-life/internal/matrix"
	"gcoletta.it/game-of-life/internal/patterns"
)

var editorPatterns = [...]patterns.Pattern{
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
	callback         Callbacks
}

func (e *editor) iteratePattern() {
	e.editorPatternIdx++
	if e.editorPatternIdx >= len(editorPatterns) {
		e.editorPatternIdx = 0
	}
}

func (e *editor) applyPattern(row, col int) {
	e.callback.Edit(func(mtx matrix.Matrix[bool]) matrix.Matrix[bool] {
		pattern := editorPatterns[e.editorPatternIdx]
		if pattern == nil {
			mtx[row][col] = !mtx[row][col]
		}
		b := matrix.Matrix[bool](mtx)
		p := matrix.Matrix[bool](pattern)
		matrix.Write(p, &b, row, col)
		return mtx
	})
}

func (e *editor) currentPattern() patterns.Pattern {
	return editorPatterns[e.editorPatternIdx]
}
