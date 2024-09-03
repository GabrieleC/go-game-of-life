package conwayadapter

import (
	"gcoletta.it/game-of-life/internal/conwaylogic"
	"gcoletta.it/game-of-life/internal/game"
)

func Iterate(old game.Matrix) game.Matrix {
	iterated := conwaylogic.Iterate(conwaylogic.Matrix(old))
	return game.Matrix(iterated)
}
