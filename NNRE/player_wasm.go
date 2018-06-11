// +build wasm

package main

import (
	"fmt"
	"syscall/js"
)

type mycb struct {
	id   string
	cell int
	p    *Player
}

func (m mycb) cb([]js.Value) {
	move := m.cell
	if move == 10 {
		m.p.inputMove <- move
	}
	if move > 9 {
		return
	}
	if move != 9 {
		if m.p.board[move] == 1 || m.p.board[move+9] == 1 {
			return
		}
		m.p.board[move] = 1
	}

	js.Global.Get("document").Call("getElementById", m.id).Set("innerHTML", js.ValueOf("O"))
	m.p.visualBoard[move] = "O"
	m.p.inputMove <- move
	m.p.hasPlayed <- "O"
	myMove := <-m.p.predictedMove
	m.p.inputMove <- myMove
	js.Global.Get("document").Call("getElementById", fmt.Sprintf("ttt%v", myMove)).Set("innerHTML", js.ValueOf("X"))
	m.p.hasPlayed <- "X"

}

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
	// Reset the content and set the handler
	for i, v := range []string{"ttt0", "ttt1", "ttt2", "ttt3", "ttt4", "ttt5", "ttt6", "ttt7", "ttt8"} {
		m := mycb{
			v,
			i,
			p,
		}
		js.Global.Get("document").Call("getElementById", v).Call("addEventListener", "click", js.NewCallback(m.cb))

	}
	go func() {
		inputMove <- 9
		p.hasPlayed <- "O"
		myMove := <-predictedMove
		js.Global.Set("output", js.ValueOf(myMove))
		js.Global.Get("document").Call("getElementById", fmt.Sprintf("ttt%v", myMove)).Set("innerHTML", js.ValueOf("X"))
		inputMove <- myMove
		p.hasPlayed <- "X"
	}()
	var cb js.Callback
	cb = js.NewCallback(func(args []js.Value) {
		move := args[0].Int()
		fmt.Println(move)
		if move == 10 {
			inputMove <- move
		}
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
