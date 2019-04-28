package main

import (
	"github.com/wangyuntao/term"
	"github.com/wangyuntao/terminfo"
)

func (d *dl) printDir(fd *fd, selected bool, id rune, printID bool) error {
	ti := term.Terminfo()
	dir := d.dirs[fd.i]

	if printID {
		color := terminfo.ColorDefault

		maxColor, ok := ti.GetNum(terminfo.MaxColors)
		if ok && maxColor >= 256 {
			color = int(id) + 21
		}

		err := ti.Text().ColorFg(color).Printf("%c ", id)
		if err != nil {
			return err
		}
		//		term.Print(id, " ")
	} else {
		err := term.Print("  ")
		if err != nil {
			return err
		}
	}

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
