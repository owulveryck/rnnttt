package main

import tm "github.com/buger/goterm"

type board [9]string

func (b board) draw() {
	tm.Clear() // Clear current screen

	tm.MoveCursor(1, 1)
	tm.Println("   |   |\n---+---+---\n   |   |\n---+---+---\n   |   |\n")
	for i, v := range b {

		switch i {
		case 0:
			tm.MoveCursor(2, 1)
		case 1:
			tm.MoveCursor(6, 1)
		case 2:
			tm.MoveCursor(10, 1)
		case 3:
			tm.MoveCursor(2, 3)
		case 4:
			tm.MoveCursor(6, 3)
		case 5:
			tm.MoveCursor(10, 3)
		case 6:
			tm.MoveCursor(2, 5)
		case 7:
			tm.MoveCursor(6, 5)
		case 8:
			tm.MoveCursor(10, 5)
		default:
		}
		tm.Println(v)
	}
	tm.MoveCursor(1, 5)
	tm.Flush() // Call it every time at the end of rendering

}
