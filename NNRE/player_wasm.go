// +build wasm

package main

import (
	"syscall/js"
)

// NewPlayer ...
func NewPlayer() *Player {
	inputMove := make(chan int, 0)
	predictedMove := make(chan int, 0)
	p := &Player{
		inputMove:     inputMove,
		predictedMove: predictedMove,
		board:         make([]int, 18),
	}

	// Dummy for imports
	_ = js.Global
	// TODO
	return p
}
