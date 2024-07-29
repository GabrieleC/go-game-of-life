package d2dui

import (
	"runtime"

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
	game          game.Game
	matrix        game.Matrix
	matrixwin     *matrixwin.Window
	redraw        bool
	width, height int
	stopRequested bool
}

func New() *D2dui {
	width := 800
	height := 800
	return &D2dui{
		redraw:    true,
		width:     width,
		height:    height,
		matrixwin: &matrixwin.Window{},
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

func (ui *D2dui) SetGame(game game.Game) {
	ui.game = game
}

func (ui *D2dui) UpdateMatrix(update game.MatrixUpdater) {
	newMatrix := update(ui.matrix)
	if ui.matrix == nil {
		ui.updateWindow()
	}
	ui.matrix = newMatrix
	ui.invalidate()
}

func (ui *D2dui) invalidate() {
	ui.redraw = true
}

func (ui *D2dui) display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gc := draw2dgl.NewGraphicContext(ui.width, ui.height)
	ui.drawGrid(gc)
	gl.Flush()
}

func (ui *D2dui) drawGrid(gc *draw2dgl.GraphicContext) {
	rows, cols := ui.matrixwin.Dimensions()
	width, height := ui.calculateGridDimensions()

	gridDesc := grid.Description{
		OriginX: 0,
		OriginY: 0,
		Width:   width,
		Height:  height,
		Rows:    rows,
		Cols:    cols,
	}

	grid.Draw(gc, gridDesc)
	ui.drawAliveCells(gc, gridDesc)
}

func (ui *D2dui) drawAliveCells(gc *draw2dgl.GraphicContext, gridDesc grid.Description) {

	maxRow, maxCol := ui.matrixwin.Dimensions()

	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			matrixRow, matrixCol := ui.matrixwin.Coords(row, col)
			if ui.matrix[matrixRow][matrixCol] {
				grid.DrawAliveCell(gc, gridDesc, row, col)
			}
		}
	}

}

func (ui *D2dui) calculateGridDimensions() (int, int) {
	cols, rows := ui.matrixwin.Dimensions()
	var width, height int

	tallerThanWiderComparision := isTallerThanWider(
		area{float64(ui.width), float64(ui.height)},
		area{float64(cols), float64(rows)},
	)

	if tallerThanWiderComparision > 0 {
		matrixRatio := float64(cols) / float64(rows)
		width = int(float64(ui.height) * matrixRatio)
		height = ui.height
	} else {
		matrixRatio := float64(rows) / float64(cols)
		height = int(float64(ui.width) * matrixRatio)
		width = ui.width
	}

	return width, height
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

	ui.updateWindow()
	ui.invalidate()
}

func (ui *D2dui) updateWindow() {
	ui.matrixwin.Update(ui.matrix.Rows(), ui.matrix.Cols())
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
			ui.game.SpeedDown()
		case key == glfw.KeyM:
			ui.game.SpeedUp()
		case key == glfw.KeyK:
			ui.game.Back()
		case key == glfw.KeyL:
			ui.game.Next()
		case key == glfw.KeyO:
			ui.matrixwin.ZoomIn()
			ui.invalidate()
		case key == glfw.KeyP:
			ui.matrixwin.ZoomOut()
			ui.invalidate()
		}
	}
}
