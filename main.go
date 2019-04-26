package main

import (
	"flag"
	"fmt"

	"github.com/wangyuntao/term"
)

var (
	sdir string // TODO
)

func main() {
	flag.Parse()
	err := term.Do(newDl().loop)
	if err != nil {
		panic(err)
	}
	fmt.Println(sdir)
}
