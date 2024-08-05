package game

import (
	"gcoletta.it/game-of-life/internal/game/periodicjob"
)

type MatrixUpdater func(old Matrix) Matrix

type UserInterface interface {
	Start() error
	Stop()
	SetGame(callback Game)
	UpdateMatrix(matrix Matrix)
}

type Game interface {
	Start() error
	Quit()
	Play()
	Pause()
	TogglePlayPause()
	SpeedUp()
	SpeedDown()
	Back()
	Next()
}

type Options struct {
	Fps           int
	InitialMatrix Matrix
}

type GameImpl struct {
	ui              UserInterface
	history         History
	fps             int
	currentlyPlay   bool
	updateChan      chan MatrixUpdater
	forwardJob      periodicjob.PeriodicJob
	listenerChannel chan struct{}
}

func New(ui UserInterface, opts Options) *GameImpl {

	history := History{timeline: make([]Matrix, 0, 100)}

	if opts.InitialMatrix != nil {
		history.append(opts.InitialMatrix)
		ui.UpdateMatrix(opts.InitialMatrix)
	}

	fps := 1
	if opts.Fps > 0 {
		fps = opts.Fps
	}

	game := GameImpl{
		ui:              ui,
		fps:             fps,
		history:         history,
		updateChan:      make(chan MatrixUpdater),
		listenerChannel: make(chan struct{}),
	}

	ui.SetGame(&game)

	go game.listenUpdates()
	game.forwardJob = periodicjob.New(fpsToInterval(fps), game.periodicForward)

	return &game
}

func (game *GameImpl) Start() error {
	err := game.ui.Start()
	return err
}

func (game *GameImpl) Play() {
	game.currentlyPlay = true
}

func (game *GameImpl) Pause() {
	game.currentlyPlay = false
}

func (game *GameImpl) TogglePlayPause() {
	if game.currentlyPlay {
		game.Pause()
	} else {
		game.Play()
	}
}

func (game *GameImpl) SpeedUp() {
	game.updateFps(game.fps + 1)
}

func (game *GameImpl) SpeedDown() {
	if game.fps > 1 {
		game.updateFps(game.fps - 1)
	}
}

func (game *GameImpl) Quit() {
	game.ui.Stop()
	game.forwardJob.Cancel()
	if game.listenerChannel != nil {
		close(game.listenerChannel)
		game.listenerChannel = nil
	}
}

func (game *GameImpl) Back() {
	game.Pause()

	matrix := game.history.back()
	if matrix != nil {
		game.ui.UpdateMatrix(matrix)
	}
}

func (game *GameImpl) Next() {
	game.Pause()
	game.forward()
}

func (game *GameImpl) UpdateMatrix(update MatrixUpdater) {
	game.updateChan <- update
}

func (game *GameImpl) updateFps(fps int) {
	game.fps = fps
	game.forwardJob.SetInterval(fpsToInterval(game.fps))
}

func (game *GameImpl) forward() {
	matrix := game.history.forward()
	if matrix != nil {
		game.ui.UpdateMatrix(matrix)
	} else {
		game.UpdateMatrix(Iterate)
	}
}

func (game *GameImpl) periodicForward() {
	if game.currentlyPlay {
		game.forward()
	}
}

func (game *GameImpl) listenUpdates() {
	for {
		select {
		case update := <-game.updateChan:
			game.applyUpdate(update)
		case <-game.listenerChannel:
			break
		}
	}
}

func (game *GameImpl) applyUpdate(update MatrixUpdater) {
	matrix := update(game.history.peek())
	game.history.append(matrix)
	game.history.forward()
	game.ui.UpdateMatrix(matrix)
}

func fpsToInterval(fps int) int {
	return 1_000_000_000 / fps
}
