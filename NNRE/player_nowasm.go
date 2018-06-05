// +build !wasm

package main

import (
	"fmt"
	"strconv"
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
	p.visualBoard.draw()

	go func() {
		inputMove <- 9
		p.play = false
		inputMove <- <-predictedMove
		for {
			p.play = true
			var err error
			fmt.Print("Enter move: ")
			var input string
			fmt.Scanln(&input)
			move, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if move > 9 {
				continue
			}
			if move != 9 {
				if p.board[move] == 1 || p.board[move+9] == 1 {
					continue
				}
				p.board[move] = 1
			}
			p.visualBoard[move] = "O"
			inputMove <- move
			p.play = false
			inputMove <- <-predictedMove
		}
	}()
	return p
}
