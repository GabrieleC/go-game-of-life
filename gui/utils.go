package gui

func initMatrix(rows, cols int) Matrix {

	matrix := make([][]bool, rows)
	for i := range matrix {
		matrix[i] = make([]bool, cols)
	}

	return matrix
}

func coalesce[T any](value *T, def T) T {
	if (value != nil) {
		return *value
	} else {
		return def
	}
}
