package main

import (
	"flag"
	"fmt"

	"github.com/wangyuntao/term"
)

var parentDir bool

func init() {
	flag.BoolVar(&parentDir, "p", false, "search parent directories")
}

func main() {
	flag.Parse()

	dl := newDl()
	err := term.Do(dl.loop)
	if err != nil {
		panic(err)
	}
	fmt.Println(dl.s)
}
