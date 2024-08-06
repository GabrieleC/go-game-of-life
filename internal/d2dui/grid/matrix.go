package grid

type Matrix [][]byte

const (
	Dead = iota
	Alive
	Shadow
)

func dimension(matrix Matrix) (rows, cols int) {
	rows = len(matrix)
	if rows > 0 {
		cols = len(matrix[0])
	}
	return rows, cols
}