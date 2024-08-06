package d2dui

import (
	"runtime"

	"gcoletta.it/game-of-life/internal/d2dui/area"
	"gcoletta.it/game-of-life/internal/d2dui/grid"
	"gcoletta.it/game-of-life/internal/d2dui/matrixwin"
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
	matrixwin     matrixwin.Matrixwin
	redraw        bool
	width, height int
	stopRequested bool
}

func New(width, height int) *D2dui {
	return &D2dui{
		redraw: true,
		width:  width,
		height: height,
	}
}

func (ui *D2dui) Start() error {
	err := glfw.Init()
	if err != nil {
		return err
	}

	window, err := glfw.CreateWindow(ui.width, ui.height, "Show RoundedRect", nil, nil)
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

	ui.reshape(window, ui.width, ui.height)

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

func (ui *D2dui) SetCallback(callback game.UICallback) {
	ui.callback = callback
}

func (ui *D2dui) UpdateMatrix(matrix game.Matrix) {
	ui.matrix = matrix
	ui.matrixwin.Update(matrix.Rows(), matrix.Cols())
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

func (ui *D2dui) createGridMatrix() grid.Matrix {
	rows, cols := ui.matrixwin.Dimensions()
	originRow, originCol := ui.matrixwin.Origin()

	// crea matrice nuova
	mtx := make([][]byte, rows)
	for rowId := range ui.matrix {
		mtx[rowId] = make([]byte, cols)
	}

	// TODO: popola matrice con celle ombra

	// popola matrice con celle vive
	for rowId, row := range ui.matrix {
		for colId := range row {
			if ui.matrix[rowId][colId] {
				mtx[rowId+originRow][colId+originCol] = grid.Alive
			}
		}
	}

	return mtx
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
		}
	}

	if action == glfw.Press || action == glfw.Repeat {
		switch {
		case key == glfw.KeyRight:
			ui.matrixwin.HorizontalPan(1)
			ui.invalidate()
		case key == glfw.KeyLeft:
			ui.matrixwin.HorizontalPan(-1)
			ui.invalidate()
		case key == glfw.KeyUp:
			ui.matrixwin.VerticalPan(-1)
			ui.invalidate()
		case key == glfw.KeyDown:
			ui.matrixwin.VerticalPan(1)
			ui.invalidate()
		case key == glfw.KeyN:
			ui.callback.SpeedDown()
		case key == glfw.KeyM:
			ui.callback.SpeedUp()
		case key == glfw.KeyK:
			ui.callback.Back()
		case key == glfw.KeyL:
			ui.callback.Next()
		case key == glfw.KeyO:
			ui.matrixwin.ZoomIn()
			ui.invalidate()
		case key == glfw.KeyP:
			ui.matrixwin.ZoomOut()
			ui.invalidate()
		}
	}
}
