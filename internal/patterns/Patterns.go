package patterns

import (
	_ "embed"

	"gcoletta.it/game-of-life/internal/matrix"
)

//go:embed txt/glider.txt
var glider string

//go:embed txt/pulsar.txt
var pulsar string

//go:embed txt/block.txt
var block string

//go:embed txt/lwss.txt
var lwss string

//go:embed txt/mwss.txt
var mwss string

//go:embed txt/hwss.txt
var hwss string

//go:embed txt/gosper.txt
var gosper string

type Writer func(mtx Pattern, row, col int)

type Pattern matrix.Matrix[bool]

func Glider() Pattern {
	return unmarshal(glider)
}

func Pulsar() Pattern {
	return unmarshal(pulsar)
}

func Block() Pattern {
	return unmarshal(block)
}

func LWSS() Pattern {
	return unmarshal(lwss)
}

func MWSS() Pattern {
	return unmarshal(mwss)
}

func HWSS() Pattern {
	return unmarshal(hwss)
}

func Gosper() Pattern {
	return unmarshal(gosper)
}
