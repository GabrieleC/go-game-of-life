package adapters

import (
	"gcoletta.it/game-of-life/internal/d2dui"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/matrix"
)

type d2duiAdapter struct {
	impl     *d2dui.D2dui
	callback game.UICallback
}

func D2duiAdapter(impl *d2dui.D2dui) game.UserInterface {
	return d2duiAdapter{impl: impl}
}

func (a d2duiAdapter) Start() error {
	return a.impl.Start()
}

func (a d2duiAdapter) Stop() {
	a.impl.Stop()
}

func (a d2duiAdapter) SetCallback(callback game.UICallback) {
	a.callback = callback
	a.impl.SetCallback(a)
}

func (a d2duiAdapter) UpdateMatrix(matrix game.Matrix) {
	x := makeGridMatrix(matrix)
	a.impl.UpdateMatrix(x)
}

func (a d2duiAdapter) Quit()            { a.callback.Quit() }
func (a d2duiAdapter) Play()            { a.callback.Play() }
func (a d2duiAdapter) Pause()           { a.callback.Pause() }
func (a d2duiAdapter) TogglePlayPause() { a.callback.TogglePlayPause() }
func (a d2duiAdapter) SpeedUp()         { a.callback.SpeedUp() }
func (a d2duiAdapter) SpeedDown()       { a.callback.SpeedDown() }
func (a d2duiAdapter) Back()            { a.callback.Back() }
func (a d2duiAdapter) Next()            { a.callback.Next() }

func (a d2duiAdapter) Edit(updater d2dui.MatrixUpdater) {
	a.callback.Edit(func(old game.Matrix) game.Matrix {
		return game.Matrix(updater(matrix.Matrix[bool](old)))
	})
}
