package main

func (d *dl) runeAdd(r rune) {
	if d.cursor == len(d.rs) {
		d.rs = append(d.rs, r)
	} else {
		rs := d.rs
		d.rs = make([]rune, len(rs)+1)
		copy(d.rs[:d.cursor], rs[:d.cursor])
		d.rs[d.cursor] = r
		copy(d.rs[d.cursor+1:], rs[d.cursor:])
	}
	d.cursor++
}

func (d *dl) runeDelB() {
	if d.cursor > 0 {
		rs := d.rs
		cursor := d.cursor
		d.cursor--
		d.rs = rs[:d.cursor]
		if cursor < len(rs) {
			d.rs = append(d.rs, rs[cursor:]...)
		}
	}
}

func (d *dl) runeDelF() {
	if d.cursor < len(d.rs) {
		rs := d.rs
		d.rs = rs[:d.cursor]
		if d.cursor+1 < len(rs) {
			d.rs = append(d.rs, rs[d.cursor+1:]...)
		}
	}
}

func (d *dl) runeDelEol() {
	d.rs = d.rs[:d.cursor]
}

func (d *dl) cursorF() {
	d.cursor++
	if d.cursor > len(d.rs) {
		d.cursor = len(d.rs)
	}
}

func (d *dl) cursorB() {
	d.cursor--
	if d.cursor < 0 {
		d.cursor = 0
	}
}

func (d *dl) cursorEol() {
	d.cursor = len(d.rs)
}

func (d *dl) cursorBol() {
	d.cursor = 0
}
