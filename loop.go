package main

import (
	"github.com/wangyuntao/term"
)

func (d *dl) loop() error {
	err := d.collect()
	if err != nil {
		return err
	}

	term.EnableWriteBuf()
	defer term.DisableWriteBuf()

	for {
		err = d.refresh()
		if err != nil {
			return err
		}

		e := term.PollEvent()

		switch v := e.(type) {
		case term.Key:
			switch v {
			case term.KeyCtrlF:
				d.cursorF()
			case term.KeyCtrlB:
				d.cursorB()
			case term.KeyCtrlE:
				d.cursorEol()
			case term.KeyCtrlA:
				d.cursorBol()
			case term.KeyBackspace:
				d.runeDelB()
			case term.KeyCtrlD:
				d.runeDelF()
			case term.KeyCtrlK:
				d.runeDelEol()
			case term.KeyCtrlP:
				d.prev()
			case term.KeyCtrlN:
				d.next()
			case term.KeyEnter:
				sdir = d.current()
				return nil

			}
		case term.Rune:
			d.runeAdd(v)
		}
	}
}
