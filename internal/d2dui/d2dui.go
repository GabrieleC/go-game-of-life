package d2dui

import (
	"runtime"

	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/geometry"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/llgcode/draw2d/draw2dgl"
)

func init() {
	runtime.LockOSThread()
}

const panAmount = 5
const cellMaxSize = 100

type D2dui struct {
	callback      game.UICallback
	matrix        game.Matrix
	origx, origy  int
	cellSize      int
	redraw        bool
	width, height int
	stopRequested bool
	curX, curY    float64
	editor        *editor
}

func New(width, height int) *D2dui {
	return &D2dui{
		redraw:   true,
		width:    width,
		height:   height,
		cellSize: 30,
		editor:   &editor{},
	}
}

func (ui *D2dui) Start() error {
	window, err := ui.initWindow()
	if err != nil {
		return err
	}

	ui.loop(window)
	return nil
}

func (ui *D2dui) initWindow() (*glfw.Window, error) {
	err := glfw.Init()
	if err != nil {
		return nil, err
	}

	window, err := glfw.CreateWindow(ui.width, ui.height, "Go - Game Of Life", nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(ui.reshape)
	window.SetKeyCallback(ui.onKey)
	window.SetCursorPosCallback(ui.onCursorPos)
	window.SetMouseButtonCallback(ui.onMouseButton)
	window.SetCloseCallback(ui.onWindowClose)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		return nil, err
	}

	ui.reshape(window, ui.width, ui.height)
	return window, nil
}

func (ui *D2dui) loop(window *glfw.Window) {
	for !ui.stopRequested {
		if ui.redraw {
			ui.display()
			window.SwapBuffers()
			ui.redraw = false
		}
		glfw.PollEvents()
	}
}

func (ui *D2dui) Stop() {
	ui.stopRequested = true
	glfw.Terminate()
}

func (ui *D2dui) SetCallback(callback game.UICallback) {
	ui.callback = callback
	ui.editor.callback = callback
}

func (ui *D2dui) UpdateMatrix(matrix game.Matrix) {
	ui.matrix = matrix
	ui.invalidate()
}

func (ui *D2dui) invalidate() {
	ui.redraw = true
}

func (ui *D2dui) display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gc := draw2dgl.NewGraphicContext(ui.width, ui.height)
	gridArea := geometry.Area{Width: ui.width, Height: ui.height}
	mtx := ui.createGridMatrix()
	origin := geometry.Point{X: ui.origx, Y: ui.origy}
	grid.Draw(gc, mtx, gridArea, origin, ui.cellSize)
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
	ui.width, ui.height = w, h
	ui.invalidate()
}

func (ui *D2dui) onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		switch {
		case (key == glfw.KeyEscape || key == glfw.KeyQ):
			ui.callback.Quit()
		case key == glfw.KeySpace:
			ui.callback.TogglePlayPause()
		case key == glfw.KeyP:
			ui.editor.iteratePattern()
			ui.invalidate()
		}
	}

	if action == glfw.Press || action == glfw.Repeat {
		switch {
		case key == glfw.KeyDown:
			ui.callback.SpeedDown()
		case key == glfw.KeyUp:
			ui.callback.SpeedUp()
		case key == glfw.KeyLeft:
			ui.callback.Back()
		case key == glfw.KeyRight:
			ui.callback.Next()
		case key == glfw.KeyW:
			ui.origy = max(ui.origy-panAmount, 0)
			ui.invalidate()
		case key == glfw.KeyS:
			maxy := (ui.cellSize * ui.matrix.Rows()) - ui.height
			ui.origy = min(ui.origy+panAmount, maxy)
			ui.invalidate()
		case key == glfw.KeyA:
			ui.origx = max(ui.origx-panAmount, 0)
			ui.invalidate()
		case key == glfw.KeyD:
			maxx := (ui.cellSize * ui.matrix.Cols()) - ui.width
			ui.origx = min(ui.origx+panAmount, maxx)
			ui.invalidate()
		case key == glfw.KeyM:
			ui.cellSize = max(ui.cellSize-1, 3)
			ui.invalidate()
		case key == glfw.KeyN:
			ui.cellSize = min(ui.cellSize+1, cellMaxSize)
			ui.invalidate()
		}
	}
}

func (ui *D2dui) onCursorPos(w *glfw.Window, xpos float64, ypos float64) {
	ui.curX = xpos
	ui.curY = ypos
	ui.invalidate()
}

func (ui *D2dui) onMouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if button == glfw.MouseButton1 && action == glfw.Release {
		ui.onLeftClick()
	}
}

func (ui *D2dui) onLeftClick() {
	gridArea := geometry.Area{Width: ui.width, Height: ui.height}
	origin := geometry.Point{X: ui.origx, Y: ui.origy}
	row, col, ok := grid.CanvasCoords(ui.curX, ui.curY, gridArea, ui.matrix.Rows(), ui.matrix.Cols(), origin, ui.cellSize)
	if ok {
		ui.editor.applyPattern(row, col)
	}
}

func (ui *D2dui) onWindowClose(w *glfw.Window) {
	ui.callback.Quit()
}
