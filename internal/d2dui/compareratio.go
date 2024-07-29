package d2dui

type area struct {
	width, height float64
}

// Returns 1 if a and b have the same ratio.
// Returns < 0 when a is "taller than wider" than b, > 0 otherwise
func isTallerThanWider(a, b area) float64 {
	aRatio := a.height / a.width
	bRatio := b.height / b.width
	return bRatio - aRatio
}
