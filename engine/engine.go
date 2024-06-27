package engine

import "gcoletta.it/game-of-life/utils"

type Grid [][]bool

type coord struct {
	row int
	col int
}

func Iterate(g Grid) Grid {
	newGrid := utils.InitMatrix(len(g), len(g[0]))



	return newGrid
}

func countActiveNeighbours(g Grid, row, col int) int {
	cells := neighbourCells(row, col)

	var count = 0
	for i := 0; i < 8; i++ {
		c := cells[i]
		if coordExists(g, c) {
			if g[c.row][c.col] {
				count++
			}
		}
	}

	return count
}

func neighbourCells(row, col int) [8]coord {
	return [8]coord{
		{row - 1, col - 1},
		{row - 1, col},
		{row - 1, col + 1},
		{row, col - 1},
		{row, col + 1},
		{row + 1, col - 1},
		{row + 1, col},
		{row + 1, col + 1},
	}
}

func coordExists(g Grid, c coord) bool {
	return rowExists(g, c.row) && colExists(g, c.col)
}

func rowExists(g Grid, row int) bool {
	return row >= 0 && row < len(g)
}

func colExists(g Grid, col int) bool {
	return col >= 0 && col < len(g[0])
}
