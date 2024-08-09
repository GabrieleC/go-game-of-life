package game

type MatrixUpdater func(old Matrix) Matrix

type UserInterface interface {
	Start() error
	Stop()
	SetCallback(callback UICallback)
	UpdateMatrix(matrix Matrix)
}

type UICallback interface {
	Quit()
	Play()
	Pause()
	TogglePlayPause()
	SpeedUp()
	SpeedDown()
	Back()
	Next()
	Edit(updater MatrixUpdater)
}
