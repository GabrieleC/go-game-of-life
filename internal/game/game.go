package game

import (
	"time"

	"gcoletta.it/game-of-life/internal/periodicjob"
)

type Game interface {
	UICallback
	Start() error
}
type Options struct {
	Fps           int
	InitialMatrix Matrix
}
type Impl struct {
	ui            UserInterface
	logic         GameLogic
	history       History
	fps           int
	currentlyPlay bool
	updateChan    chan MatrixUpdater
	forwardJob    periodicjob.PeriodicJob
	quitChan      chan struct{}
}

func New(ui UserInterface, logic GameLogic, opts Options) *Impl {

	history := History{timeline: make([]Matrix, 0, 100)}

	if opts.InitialMatrix != nil {
		history.append(opts.InitialMatrix)
		ui.UpdateMatrix(opts.InitialMatrix)
	}

	fps := 1
	if opts.Fps > 0 {
		fps = opts.Fps
	}

	game := Impl{
		ui:         ui,
		logic:      logic,
		fps:        fps,
		history:    history,
		updateChan: make(chan MatrixUpdater),
		quitChan:   make(chan struct{}),
	}

	ui.SetCallback(&game)
	go game.listenUpdates()
	game.forwardJob = periodicjob.New(fpsToDuration(fps), game.periodicForward)

	return &game
}

func (game *Impl) Start() error {
	err := game.ui.Start()
	return err
}

func (game *Impl) Play() {
	game.currentlyPlay = true
}

func (game *Impl) Pause() {
	game.currentlyPlay = false
}

func (game *Impl) TogglePlayPause() {
	if game.currentlyPlay {
		game.Pause()
	} else {
		game.Play()
	}
}

func (game *Impl) SpeedUp() {
	game.updateFps(game.fps + 1)
}

func (game *Impl) SpeedDown() {
	if game.fps > 1 {
		game.updateFps(game.fps - 1)
	}
}

func (game *Impl) Quit() {
	game.ui.Stop()
	game.forwardJob.Cancel()
	if game.quitChan != nil {
		close(game.quitChan)
		game.quitChan = nil
	}
}

func (game *Impl) Back() {
	game.Pause()

	matrix := game.history.back()
	if matrix != nil {
		game.ui.UpdateMatrix(matrix)
	}
}

func (game *Impl) Next() {
	game.Pause()
	game.forward()
}

func (game *Impl) Edit(updater MatrixUpdater) {
	game.updateMatrix(updater)
}

func (game *Impl) updateMatrix(update MatrixUpdater) {
	game.updateChan <- update
}

func (game *Impl) updateFps(fps int) {
	game.fps = fps
	game.forwardJob.SetInterval(fpsToDuration(game.fps))
}

func (game *Impl) forward() {
	matrix := game.history.forward()
	if matrix != nil {
		game.ui.UpdateMatrix(matrix)
	} else {
		game.updateMatrix(MatrixUpdater(game.logic))
	}
}

func (game *Impl) periodicForward() {
	if game.currentlyPlay {
		game.forward()
	}
}

func (game *Impl) listenUpdates() {
	for {
		select {
		case update := <-game.updateChan:
			game.applyUpdate(update)
		case <-game.quitChan:
			break
		}
	}
}

func (game *Impl) applyUpdate(update MatrixUpdater) {
	matrix := update(game.history.peek())
	game.history.append(matrix)
	game.history.forward()
	game.ui.UpdateMatrix(matrix)
}

func fpsToDuration(fps int) time.Duration {
	nanos := int(time.Second.Nanoseconds()) / fps
	return time.Nanosecond * time.Duration(nanos)
}
