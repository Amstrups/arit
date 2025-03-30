package tui

type Square struct {
	X, Y       int
	W, H       int
	A, B, C, D int
}

type Section struct {
	outer, inner Square
}
