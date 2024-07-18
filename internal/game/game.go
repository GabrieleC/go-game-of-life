package game

import (
	"time"
)

type Game struct {
	gui           Gui
	fps           int
	altRequested  bool
	currentlyPlay bool
}

type MatrixUpdater func(old Matrix) Matrix

type Callbacks struct {
	Quit            func()
	Pause           func()
	Play            func()
	TogglePlayPause func()
	SpeedUp         func()
	SpeedDown       func()
}

type Gui interface {
	Start() error
	Stop()
	SetCallbacks(callbacks Callbacks)
	UpdateMatrix(update MatrixUpdater)
}

type Options struct {
	Fps           int
	InitialMatrix Matrix
}

func New(gui Gui, opts Options) Game {

	if opts.InitialMatrix != nil {
		gui.UpdateMatrix(func(m Matrix) Matrix {
			return opts.InitialMatrix
		})
	}

	fps := 1
	if opts.Fps > 0 {
		fps = opts.Fps
	}

	return Game{gui: gui, fps: fps}
}

func (game *Game) Execute() {

	game.gui.SetCallbacks(Callbacks{
		Quit:            game.quit,
		SpeedUp:         game.speedUp,
		SpeedDown:       game.speedDown,
		TogglePlayPause: game.TogglePlayPause,
	})

	game.Play()
}

func (game *Game) Play() {
	game.currentlyPlay = true
	go func() {
		for {
			if game.altRequested {
				game.altRequested = false
				break
			}
			time.Sleep(time.Duration(1_000_000_000 / game.fps))
			game.gui.UpdateMatrix(Iterate)
		}
	}()
}

func (game *Game) Pause() {
	game.currentlyPlay = false
	game.altRequested = true
}

func (game *Game) TogglePlayPause() {
	if game.currentlyPlay {
		game.Pause()
	} else {
		game.Play()
	}
}

func (game *Game) quit() {
	game.altRequested = true
	game.gui.Stop()
}

func (game *Game) speedUp() {
	game.fps += 1
}

func (game *Game) speedDown() {
	if game.fps > 1 {
		game.fps -= 1
	}
}
