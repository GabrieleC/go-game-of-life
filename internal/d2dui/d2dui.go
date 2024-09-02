package d2dui

import (
	"runtime"

	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/game"
	"gcoletta.it/game-of-life/internal/geom"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/llgcode/draw2d/draw2dgl"
)

const defaultWinSize = 840
const defaultTitle = "Go - Game Of Life"

func init() {
	runtime.LockOSThread()
}

type Callbacks interface {
	Quit()
	Play()
	Pause()
	TogglePlayPause()
	SpeedUp()
	SpeedDown()
	Back()
	Next()
	Edit(updater game.MatrixUpdater) // TODO
}

type D2dui struct {
	width, height int
	matrix        grid.Matrix
	callback      Callbacks
	redraw        bool
	stopRequested bool
	grd           grid.Grid
	cursor        geom.Point
	editor        *editor
	gc            *draw2dgl.GraphicContext
}

func New(width, height int) *D2dui {

	g := grid.Grid{
		Origin:   geom.Point{},
		CellSize: 30,
		Canvas:   geom.Area{Width: width, Height: height},
	}

	return &D2dui{
		redraw: true,
		width:  width,
		height: height,
		editor: &editor{},
		grd:    g,
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

	window, err := glfw.CreateWindow(defaultWinSize, defaultWinSize, defaultTitle, nil, nil)
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

func (ui *D2dui) SetCallback(callback Callbacks) {
	ui.callback = callback
	ui.editor.callback = callback
}

func (ui *D2dui) UpdateMatrix(matrix grid.Matrix) {
	ui.matrix = matrix
	ui.invalidate()
}

func (ui *D2dui) invalidate() {
	ui.redraw = true
}

func (ui *D2dui) display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	ui.grd.Matrix = grid.Copy(ui.matrix)
	applyEditorPattern(ui.grd, ui.grd.Matrix, ui.cursor, ui.editor.currentPattern())
	ui.grd.Draw(ui.gc)
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
	ui.gc = draw2dgl.NewGraphicContext(ui.width, ui.height)
	ui.grd.Canvas = geom.Area{Width: w, Height: h}
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
			ui.grd.PanUp()
			ui.invalidate()
		case key == glfw.KeyS:
			ui.grd.PanDown()
			ui.invalidate()
		case key == glfw.KeyA:
			ui.grd.PanLeft()
			ui.invalidate()
		case key == glfw.KeyD:
			ui.grd.PanRight()
			ui.invalidate()
		case key == glfw.KeyM:
			ui.grd.ZoomOut()
			ui.invalidate()
		case key == glfw.KeyN:
			ui.grd.ZoomIn()
			ui.invalidate()
		}
	}
}

func (ui *D2dui) onCursorPos(w *glfw.Window, xpos float64, ypos float64) {
	ui.cursor.X = int(xpos)
	ui.cursor.Y = int(ypos)
	ui.invalidate()
}

func (ui *D2dui) onMouseButton(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if button == glfw.MouseButton1 && action == glfw.Release {
		ui.onLeftClick()
	}
}

func (ui *D2dui) onLeftClick() {
	cursor, ok := ui.grd.CanvasCoords(ui.cursor)
	if ok {
		ui.editor.applyPattern(cursor.Y, cursor.X)
	}
}

func (ui *D2dui) onWindowClose(w *glfw.Window) {
	ui.callback.Quit()
}

func applyEditorPattern(grd grid.Grid, mtx grid.Matrix, cursor geom.Point, pattern game.Matrix) {
	origin, ok := grd.CanvasCoords(cursor)
	if ok {
		grid.ApplyPattern(pattern, origin, mtx, editorStateUpdater)
	}
}

func editorStateUpdater(oldState byte) byte {
	if oldState == grid.Dead {
		return grid.Shadow
	}
	return oldState
}
