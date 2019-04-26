package main

import (
	"github.com/wangyuntao/term"
	"github.com/wangyuntao/terminfo"
)

func (d *dl) refresh() error {
	if len(d.rs) != d.rsl {
		d.rsl = len(d.rs)
		d.filter()
	}

	ti := term.Terminfo()

	err := ti.ClearScreen()
	if err != nil {
		return err
	}

	row, _, err := term.WinSize()
	if err != nil {
		return err
	}

	row-- // status line
	row-- // prompt line

	if d.fdi < 0 || d.fdi >= len(d.fds) {
		d.fdi = len(d.fds) - 1
	}

	i2 := len(d.fds)
	i1 := i2 - row

	for ; i1 < 0; i1++ { // not enough lines
		err = term.Println() // write empty lines
		if err != nil {
			return err
		}
	}

	if d.fdi >= 0 && d.fdi < i1 {
		i1 = d.fdi
		i2 = i1 + row
	}

	// dirs
	for i := i1; i < i2; i++ {
		fd := d.fds[i]
		err := d.printDir(fd, i == d.fdi)
		if err != nil {
			return err
		}
	}

	// status
	err = ti.Text().ColorFg(terminfo.ColorGreen).Printf("  [%d/%d/%d]\n", d.fdi+1, len(d.fds), len(d.dirs))
	if err != nil {
		return err
	}

	// prompt
	err = term.Print("> ", string(d.rs))
	if err != nil {
		return err
	}

	err = ti.CursorAddress(row+2, 2+d.cursor)
	if err != nil {
		return err
	}

	return term.Flush()
}
