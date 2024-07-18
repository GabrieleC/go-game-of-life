package game

import (
	"time"

	"gcoletta.it/game-of-life/internal/engine"
	"gcoletta.it/game-of-life/internal/gui"
	"gcoletta.it/game-of-life/internal/matrix"
	"gcoletta.it/game-of-life/internal/patterns"
)

type Game struct {
	gui          gui.Gui
	fps          int
	altRequested bool
}

func New(gui gui.Gui, fps int) Game {
	return Game{gui: gui, fps: 1}
}

func (game *Game) Execute() {

	game.gui.SetCallbacks(gui.Callbacks{
		Quit:      game.quit,
		SpeedUp:   game.speedUp,
		SpeedDown: game.speedDown,
	})

	game.gui.UpdateMatrix(func(m gui.Matrix) gui.Matrix {
		patterns.Glider(matrix.Matrix(m), 1, 1)
		patterns.Pulsar(matrix.Matrix(m), 8, 8)
		return m
	})

	game.Play()
}

func (game *Game) Play() {
	go func() {
		for {
			if game.altRequested {
				break
			}
			time.Sleep(time.Duration(1_000_000_000 / game.fps))
			game.gui.UpdateMatrix(matrixUpdater)
		}
	}()
}

func (game *Game) Pause() {
	game.altRequested = true
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

func matrixUpdater(old gui.Matrix) gui.Matrix {
	oldGrid := engine.Grid(old)
	newGrid := engine.Iterate(oldGrid)
	return gui.Matrix(matrix.Matrix(newGrid))
}
