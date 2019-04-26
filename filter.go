package main

import (
	"regexp"
)

func (d *dl) filter() {
	var fds []*fd

	if len(d.rs) == 0 {
		fds = make([]*fd, len(d.dirs))
		for i := 0; i < len(d.dirs); i++ {
			fds[i] = &fd{i, nil}
		}
	} else {
		fds = make([]*fd, 0, 64)
		re, err := regexp.Compile(string(d.rs))
		if err != nil {
			return
		}

		for i, dir := range d.dirs {
			mis := re.FindAllStringIndex(dir, -1)
			if mis != nil {
				fds = append(fds, &fd{i, mis})
			}
		}
	}

	d.fds = fds
	d.fdi = len(fds) - 1
}
