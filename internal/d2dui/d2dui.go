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
	callback      game.UICallback
	matrix        game.Matrix
	redraw        bool
	width, height int
	stopRequested bool
	curX, curY    float64
	editor        *editor
}

func New(width, height int) *D2dui {
	return &D2dui{
		redraw: true,
		width:  width,
		height: height,
		editor: &editor{},
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
	gridArea := area.Area{Width: ui.width, Height: ui.height}
	mtx := ui.createGridMatrix()
	grid.Draw(gc, mtx, gridArea)
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
	gridArea := area.Area{Width: ui.width, Height: ui.height}
	row, col, ok := grid.CanvasCoords(ui.curX, ui.curY, gridArea, ui.matrix.Rows(), ui.matrix.Cols())
	if ok {
		ui.editor.applyPattern(row, col)
	}
}
