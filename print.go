package main

import (
	"github.com/wangyuntao/term"
	"github.com/wangyuntao/terminfo"
)

func (d *dl) printDir(fd *fd, selected bool) error {
	dir := d.dirs[fd.i]

	if selected {
		term.Print("  ")
	} else {
		term.Print("  ")
	}

	ti := term.Terminfo()

	i0 := 0
	for i := 0; i < len(fd.mis); i++ {
		rg := fd.mis[i]

		i1 := rg[0]
		i2 := rg[1]

		// s1
		s1 := dir[i0:i1]
		t1 := ti.Text()
		if selected {
			t1.Underline().ColorFg(terminfo.ColorCyan)
		}
		err := t1.Print(s1)
		if err != nil {
			return err
		}

		// s2
		s2 := dir[i1:i2]
		t2 := ti.Text()
		if selected {
			t2.Underline()
		}
		err = t2.ColorFg(terminfo.ColorYellow).Print(s2)
		if err != nil {
			return err
		}
		i0 = i2
	}

	if i0 < len(dir) {
		s3 := dir[i0:]
		t3 := ti.Text()
		if selected {
			t3.Underline().ColorFg(terminfo.ColorCyan)
		}
		err := t3.Print(s3)
		if err != nil {
			return err
		}
	}

	term.Println()
	return nil
}
