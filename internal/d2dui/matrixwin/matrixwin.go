package matrixwin

type Matrixwin struct {
	originCol, originRow int
	cols, rows           int
	maxCols, maxRows     int
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

func (w *Matrixwin) Origin() (int, int) {
	return w.originRow, w.originCol
}

func (w *Matrixwin) Coords(winY, winX int) (y, x int) {
	return winY + w.originRow, winX + w.originCol
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

		if w.cols > w.maxCols - w.originCol {
			w.originCol--
		}
		if w.rows > w.maxRows - w.originRow {
			w.originRow--
		}
	}
}

func (w *Matrixwin) HorizontalPan(steps int) {
	x := w.originCol + steps
	if x+w.cols <= w.maxCols && x >= 0 {
		w.originCol = x
	}
}

func (w *Matrixwin) VerticalPan(steps int) {
	y := w.originRow + steps
	if y+w.rows <= w.maxRows && y >= 0 {
		w.originRow = y
	}
}
