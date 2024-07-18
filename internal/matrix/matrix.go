package matrix

type Matrix [][]bool

func NewMatrix(rows, cols int) Matrix {
	matrix := make([][]bool, rows)
	for i := range matrix {
		matrix[i] = make([]bool, cols)
	}
	return matrix
}

func (destination *Matrix) Copy(source Matrix, originRow int, originCol int) {
	for rowId, row := range source {
		for colId := range row {
			(*destination)[rowId+originRow][colId+originCol] = source[rowId][colId]
		}
	}
}
