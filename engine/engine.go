package engine

import "fmt"

const ROWS = 10
const COLS = 25

type grid [ROWS][COLS]bool

type coord struct {
	row int
	col int
}

func start() {
	var g grid
	g[0][0] = true
	g[5][12] = true
	printGrid(g)
	active := countActive(g, 0, 2)
	fmt.Println(active)
}

func countActive(g grid, row, col int) int {
	cells := neighbourCells(row, col)

	var count = 0
	for i := 0; i < 8; i++ {
		c := cells[i]
		if coordExists(c) {
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

func coordExists(c coord) bool {
	return rowExists(c.row) && colExists(c.col)
}

func rowExists(row int) bool {
	return row >= 0 && row < ROWS
}

func colExists(col int) bool {
	return col >= 0 && col < COLS
}

func printGrid(g grid) {
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if g[i][j] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
