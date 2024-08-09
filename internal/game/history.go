package game
type History struct {
	timeline []Matrix
	cursor   int
}

func (history *History) append(matrix Matrix) {
	history.timeline = append(history.timeline, matrix)
}

func (history *History) peek() Matrix {
	return history.timeline[history.cursor]
}

func (history *History) back() Matrix {
	if history.cursor == 0 {
		return nil
	}
	history.cursor--
	return history.timeline[history.cursor]
}

func (history *History) forward() Matrix {
	if history.cursor + 1 >= len(history.timeline) {
		return nil
	}
	history.cursor++
	return history.timeline[history.cursor]
}
