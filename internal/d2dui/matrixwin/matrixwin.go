package matrixwin

type Matrixwin struct {
	originX, originY int
	cols, rows       int
	maxCols, maxRows int
}

func (w *Matrixwin) Update(rows, cols int) {
	w.maxCols = cols
	w.maxRows = rows

	if cols < rows {
		w.cols = cols
		w.rows = cols
	} else {
		w.cols = rows
		w.rows = rows
	}
}

func (w *Matrixwin) Dimensions() (int, int) {
	return w.rows, w.cols
}

func (w *Matrixwin) Coords(winY, winX int) (y, x int) {
	return winY + w.originY, winX + w.originX
}

func (w *Matrixwin) ZoomIn() {
	if w.cols > 1 && w.rows > 1 {
		w.cols--
		w.rows--
	}
}

func (w *Matrixwin) ZoomOut() {
	if w.cols < w.maxCols && w.rows < w.maxRows {
		w.cols++
		w.rows++
	}
}

func (w *Matrixwin) HorizontalPan(steps int) {
	x := w.originX + steps
	if x+w.cols <= w.maxCols && x >= 0 {
		w.originX = x
	}
}

func (w *Matrixwin) VerticalPan(steps int) {
	y := w.originY + steps
	if y+w.rows <= w.maxRows && y >= 0 {
		w.originY = y
	}
}
