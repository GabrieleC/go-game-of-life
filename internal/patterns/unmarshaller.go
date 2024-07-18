package patterns

import (
	"strings"

	"gcoletta.it/game-of-life/internal/game"
)

func unmarshal(input string) game.Matrix {

	rows := strings.Split(input, "\n")

	colsCount := maxLength(rows)
	rowsCount := len(rows)

	matrix := game.NewMatrix(rowsCount, colsCount)
	writeInMatrix(rows, matrix)

	return matrix
}

func writeInMatrix(rows []string, matrix [][]bool) {
	for rowId, row := range rows {
		for colId, val := range row {
			matrix[rowId][colId] = (val == 'X')
		}
	}
}

func maxLength(rows []string) int {
	maxLen := 0
	for _, v := range rows {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	return maxLen
}
