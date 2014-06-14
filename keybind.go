package main

import (
	"github.com/nsf/termbox-go"
)

func waitKeyInput() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return
			} else if ev.Ch == 'p' {
				if clock.paused {
					clock.start()
				} else {
					clock.pause()
				}
			} else if !clock.paused {
				if ev.Ch == 'z' {
					currentMino.rotateLeft()
				} else if ev.Ch == 'x' || ev.Key == termbox.KeyArrowUp {
					currentMino.rotateRight()
				} else if ev.Key == termbox.KeySpace {
					currentMino.drop()
				} else if ev.Key == termbox.KeyArrowDown {
					currentMino.moveDown()
				} else if ev.Key == termbox.KeyArrowLeft {
					currentMino.moveLeft()
				} else if ev.Key == termbox.KeyArrowRight {
					currentMino.moveRight()
				}
			}
		}
		refreshScreen()
	}
}
