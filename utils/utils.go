package utils

func InitMatrix(rows, cols int) [][]bool {

	matrix := make([][]bool, rows)
	for i := range matrix {
		matrix[i] = make([]bool, cols)
	}

	return matrix
}

func CoalescePtr[T any](value *T, def T) T {
	if value != nil {
		return *value
	} else {
		return def
	}
}

func CoalesceF32(value float32, def float32) float32 {
	if value != float32(0) {
		return value
	} else {
		return def
	}
}
