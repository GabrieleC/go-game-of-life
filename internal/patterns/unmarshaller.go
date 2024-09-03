package patterns

import (
	"strings"

	"gcoletta.it/game-of-life/internal/matrix"
)

func unmarshal(input string) Pattern {
	rows := strings.Split(input, "\n")
	colsCount := maxLength(rows)
	rowsCount := len(rows)

	matrix := Pattern(matrix.Create[bool](rowsCount, colsCount))
	writeInMatrix(rows, matrix)

	return Pattern(matrix)
}

func writeInMatrix(rows []string, pattern Pattern) {
	for rowId, row := range rows {
		for colId, val := range row {
			pattern[rowId][colId] = (val == 'X')
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
