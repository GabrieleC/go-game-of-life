package area

type Area struct {
	Width, Height int
}

// Returns 1 if a and b have the same ratio.
// Returns < 0 when a is "taller than wider" than b, > 0 otherwise
func IsTallerThanWider(a, b Area) float64 {
	aRatio := float64(a.Height) / float64(a.Width)
	bRatio := float64(b.Height) / float64(b.Width)
	return bRatio - aRatio
}
