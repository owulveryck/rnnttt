// +build wasm

package main

import (
	"fmt"
	"syscall/js"
)

// NewPlayer ...
func NewPlayer() *Player {
	fmt.Println("New player")
	inputMove := make(chan int, 0)
	predictedMove := make(chan int, 0)
	p := &Player{
		inputMove:     inputMove,
		predictedMove: predictedMove,
		board:         make([]int, 18),
		hasPlayed:     make(chan string, 0),
	}

	go func() {
		fmt.Println("Let's play")
		inputMove <- 9
		p.hasPlayed <- "O"
		myMove := <-predictedMove
		js.Global.Set("output", js.ValueOf(myMove))
		inputMove <- myMove
		p.hasPlayed <- "X"
	}()
	var cb js.Callback
	cb = js.NewCallback(func(args []js.Value) {
		fmt.Println("Callback")
		move := args[0].Int()
		fmt.Println(move)
		if move > 9 {
			return
		}
		if move != 9 {
			if p.board[move] == 1 || p.board[move+9] == 1 {
				return
			}
			p.board[move] = 1
		}

		p.visualBoard[move] = "O"
		inputMove <- move
		p.hasPlayed <- "O"
		myMove := <-predictedMove
		js.Global.Set("output", js.ValueOf(myMove))
		inputMove <- myMove
		p.hasPlayed <- "X"
	})
	js.Global.Set("play", cb)

	return p
}
