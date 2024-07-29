package game

type coord struct {
	row, col int
}

/*
	Conway's rules:
	- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
	- Any live cell with two or three live neighbours lives on to the next generation.
	- Any live cell with more than three live neighbours dies, as if by overpopulation.
	- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/

func Iterate(old Matrix) Matrix {
	new := NewMatrix(len(old), len(old[0]))
	iterateGrid(old, new)
	return new
}

func iterateGrid(oldGrid Matrix, newGrid Matrix) {
	for row := range oldGrid {
		for col := range oldGrid[row] {
			newGrid[row][col] = iterateCell(oldGrid, row, col)
		}
	}
}

func iterateCell(g Matrix, row, col int) bool {
	aliveNeighbours := countActiveNeighbours(g, row, col)
	wasAlive := g[row][col]
	return remainsAlive(wasAlive, aliveNeighbours) || comesAlive(wasAlive, aliveNeighbours)
}

func remainsAlive(wasAlive bool, aliveNeighbours int) bool {
	return wasAlive && aliveNeighbours >= 2 && aliveNeighbours <= 3
}

func comesAlive(currentAlive bool, aliveNeighbours int) bool {
	return !currentAlive && aliveNeighbours == 3
}

func countActiveNeighbours(g Matrix, row, col int) int {
	count := 0
	for _, c := range neighbourCells(row, col) {
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

func coordExists(g Matrix, c coord) bool {
	return rowExists(g, c.row) && colExists(g, c.col)
}

func rowExists(g Matrix, row int) bool {
	return row >= 0 && row < len(g)
}

func colExists(g Matrix, col int) bool {
	return col >= 0 && col < len(g[0])
}
