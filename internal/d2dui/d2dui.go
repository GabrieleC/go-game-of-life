package d2dui

import (
	"runtime"

	"gcoletta.it/game-of-life/internal/d2dui/area"
	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/llgcode/draw2d/draw2dgl"
)

func init() {
	runtime.LockOSThread()
}

type D2dui struct {
	game                game.Game
	matrix              game.Matrix
	grid                grid.Grid
	redraw              bool
	winWidth, winHeight int
	stopRequested       bool
}

func New() *D2dui {
	width := 800
	height := 800
	return &D2dui{
		redraw:    true,
		winWidth:  width,
		winHeight: height,
		grid:      grid.New(),
	}
}

func (ui *D2dui) Start() error {
	err := glfw.Init()
	if err != nil {
		return err
	}

	window, err := glfw.CreateWindow(ui.winWidth, ui.winHeight, "Show RoundedRect", nil, nil)
	if err != nil {
		return err
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(ui.reshape)
	window.SetKeyCallback(ui.onKey)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		return err
	}

	ui.reshape(window, ui.winWidth, ui.winHeight)

	for !ui.stopRequested {
		if ui.redraw {
			ui.display()
			window.SwapBuffers()
			ui.redraw = false
		}
		glfw.PollEvents()
	}

	return nil
}

func (ui *D2dui) Stop() {
	ui.stopRequested = true
	glfw.Terminate()
}

func (ui *D2dui) SetGame(game game.Game) {
	ui.game = game
}

func (ui *D2dui) UpdateMatrix(matrix game.Matrix) {
	ui.matrix = matrix
	ui.grid.UpdateMatrix(ui.matrix)
	ui.invalidate()
}

func (ui *D2dui) invalidate() {
	ui.redraw = true
}

func (ui *D2dui) display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gc := draw2dgl.NewGraphicContext(ui.winWidth, ui.winHeight)
	gridArea := area.Area{Width: ui.winWidth, Height: ui.winHeight}
	gridOrigin := grid.Coord{X: 0, Y: 0}
	ui.grid.Draw(gc, gridOrigin, gridArea)
	gl.Flush()
}

func (ui *D2dui) reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
	ui.winWidth, ui.winHeight = w, h
	ui.invalidate()
}

func (ui *D2dui) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch {
		case (key == glfw.KeyEscape || key == glfw.KeyQ):
			ui.game.Quit()
		case key == glfw.KeySpace:
			ui.game.TogglePlayPause()
		}
	}

	if action == glfw.Press || action == glfw.Repeat {
		switch {
		case key == glfw.KeyRight:
			ui.grid.HorizontalPan(1)
			ui.invalidate()
		case key == glfw.KeyLeft:
			ui.grid.HorizontalPan(-1)
			ui.invalidate()
		case key == glfw.KeyUp:
			ui.grid.VerticalPan(-1)
			ui.invalidate()
		case key == glfw.KeyDown:
			ui.grid.VerticalPan(1)
			ui.invalidate()
		case key == glfw.KeyN:
			ui.game.SpeedDown()
		case key == glfw.KeyM:
			ui.game.SpeedUp()
		case key == glfw.KeyK:
			ui.game.Back()
		case key == glfw.KeyL:
			ui.game.Next()
		case key == glfw.KeyO:
			ui.grid.ZoomIn()
			ui.invalidate()
		case key == glfw.KeyP:
			ui.grid.ZoomOut()
			ui.invalidate()
		}
	}
}
