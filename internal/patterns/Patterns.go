package patterns

import (
	_ "embed"

	"gcoletta.it/game-of-life/internal/game"
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

type Pattern func(mtx game.Matrix, row, col int)

func Glider() game.Matrix {
	return unmarshal(glider)
}

func Pulsar() game.Matrix {
	return unmarshal(pulsar)
}

func Block() game.Matrix {
	return unmarshal(block)
}

func LWSS() game.Matrix {
	return unmarshal(lwss)
}

func MWSS() game.Matrix {
	return unmarshal(mwss)
}

func HWSS() game.Matrix {
	return unmarshal(hwss)
}

func Gosper() game.Matrix {
	return unmarshal(gosper)
}
