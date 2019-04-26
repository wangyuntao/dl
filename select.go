package main

func (d *dl) prev() {
	l := len(d.fds)
	if l == 0 {
		return
	}

	d.fdi--
	if d.fdi < 0 {
		d.fdi = 0
	}
}

func (d *dl) next() {
	l := len(d.fds)
	if l == 0 {
		return
	}

	d.fdi++
	if d.fdi >= l {
		d.fdi = l - 1
	}
}

func (d *dl) current() string {
	l := len(d.fds)
	if l == 0 {
		return ""
	}

	fd := d.fds[d.fdi]
	return d.dirs[fd.i]
}
