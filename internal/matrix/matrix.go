package matrix

type Matrix[T any] [][]T

func Create[T any](rows, cols int) Matrix[T] {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}
	return matrix
}

func Write[T any](source Matrix[T], dest *Matrix[T], originRow int, originCol int) {
	for rowId, row := range source {
		for colId := range row {
			(*dest)[rowId+originRow][colId+originCol] = source[rowId][colId]
		}
	}
}

func Copy[T any](src Matrix[T]) Matrix[T] {
	duplicate := make([][]T, len(src))
	for i := range src {
		duplicate[i] = make([]T, len(src[i]))
		copy(duplicate[i], src[i])
	}
	return duplicate
}

func Dimensions[T any](mtx Matrix[T]) (rows, cols int) {
	if len(mtx) > 0 {
		return len(mtx), len(mtx[0])
	} else {
		return 0, 0
	}
}
