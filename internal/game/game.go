package game

import (
	"time"
)

type MatrixUpdater func(old Matrix) Matrix

type Gui interface {
	Start() error
	Stop()
	SetGame(callback Game)
	UpdateMatrix(update MatrixUpdater)
}

type Game interface {
	Quit()
	Pause()
	Play()
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
	gui           Gui
	history       History
	fps           int
	altRequested  bool
	currentlyPlay bool
}

func New(gui Gui, opts Options) *GameImpl {

	history := History{timeline: make([]Matrix, 0, 100)}

	if opts.InitialMatrix != nil {
		history.append(opts.InitialMatrix)
		gui.UpdateMatrix(func(m Matrix) Matrix {
			return opts.InitialMatrix
		})
	}

	fps := 1
	if opts.Fps > 0 {
		fps = opts.Fps
	}

	game := GameImpl{
		gui:     gui,
		fps:     fps,
		history: history,
	}

	return &game
}

func (game *GameImpl) Play() {
	game.currentlyPlay = true
	go func() {
		for {
			if game.altRequested {
				game.altRequested = false
				break
			}
			time.Sleep(time.Duration(1_000_000_000 / game.fps))
			game.forward()
		}
	}()
}

func (game *GameImpl) Pause() {
	if game.currentlyPlay {
		game.currentlyPlay = false
		game.altRequested = true
	}
}

func (game *GameImpl) TogglePlayPause() {
	if game.currentlyPlay {
		game.Pause()
	} else {
		game.Play()
	}
}

func (game *GameImpl) Quit() {
	game.altRequested = true
	game.gui.Stop()
}

func (game *GameImpl) SpeedUp() {
	game.fps += 1
}

func (game *GameImpl) SpeedDown() {
	if game.fps > 1 {
		game.fps -= 1
	}
}

func (game *GameImpl) Back() {
	game.Pause()

	matrix := game.history.back()
	if matrix != nil {
		game.gui.UpdateMatrix(func(old Matrix) Matrix { return matrix })
	}
}

func (game *GameImpl) Next() {
	game.Pause()
	game.forward()
}

func (game *GameImpl) forward() {
	matrix := game.history.forward()
	if matrix == nil {
		matrix = Iterate(game.history.peek())
		game.history.append(matrix)
		game.history.forward()
	}

	game.gui.UpdateMatrix(func(old Matrix) Matrix { return matrix })
}