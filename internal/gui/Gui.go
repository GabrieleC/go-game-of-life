package gui

import "gcoletta.it/game-of-life/internal/matrix"

type Matrix matrix.Matrix

type MatrixUpdater func(old Matrix) Matrix

type Callbacks struct {
	Quit func()
	Pause func()
	Play func()
	SpeedUp func()
	SpeedDown func()
}

type Gui interface {
	Start() error
	Stop()
	SetCallbacks(callbacks Callbacks)
	UpdateMatrix(update MatrixUpdater)
}
