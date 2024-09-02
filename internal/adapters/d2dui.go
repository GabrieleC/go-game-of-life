package adapters

import (
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
)

type d2duiAdapter struct {
	impl *d2dui.D2dui
}

func D2duiAdapter(impl *d2dui.D2dui) game.UserInterface {
	return d2duiAdapter{impl}
}

func (a d2duiAdapter) Start() error {
	return a.impl.Start()
}

func (a d2duiAdapter) Stop() {
	a.impl.Stop()
}

func (a d2duiAdapter) SetCallback(callback game.UICallback) {
	a.impl.SetCallback(callback)
}

func (a d2duiAdapter) UpdateMatrix(matrix game.Matrix) {
	x := makeGridMatrix(matrix)
	a.impl.UpdateMatrix(x)
}
