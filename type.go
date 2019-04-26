package main

type dl struct {
	dirs []string

	// filter & select
	fds []*fd // filtered dirs
	fdi int   // selected idx of filtered dir

	// input
	rs     []rune
	rsl    int
	cursor int
}

type fd struct {
	i   int     // dir idx
	mis [][]int // matches idxs
}

func newDl() *dl {
	return &dl{}
}
