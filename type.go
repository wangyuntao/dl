package main

type dl struct {
	cfg *cfg

	// original list of dirs
	dirs []string

	// filter & select
	fds []*fd
	fdi int

	// input
	rs     []rune
	rsl    int
	cursor int

	// result
	s string
}

type fd struct {
	i   int     // dir idx
	mis [][]int // matches idxs
}

func newDl() *dl {
	return &dl{}
}
