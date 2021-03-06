package main

import (
	"errors"
	"fmt"
	"io"
)

// Player ...
type Player struct {
	board         []int
	visualBoard   board
	inputMove     chan int
	predictedMove chan int
	offset        int
	hasPlayed     chan string
}

func (p *Player) Read() ([]float32, error) {
	mve := <-p.inputMove
	if mve == 10 {
		return nil, io.EOF
	}
	output := make([]float32, 18)
	for i := range p.board {
		output[i] = float32(p.board[i])
	}
	p.offset++
	return output, nil
}

func (p *Player) Write(v []float32) error {
	hasPlayed := <-p.hasPlayed
	if hasPlayed == "X" {
		return nil
	}
	// Get the max probability
	max := float32(0)
	idx := -1
	for i, v := range v {
		if v > max && p.board[i] == 0 && p.board[i+9] == 0 {
			max = v
			idx = i
		}
	}
	if idx == -1 {
		return errors.New("game end")
	}
	fmt.Println("My move:", idx)
	p.visualBoard[idx] = "X"
	p.visualBoard.draw()
	p.board[idx+9] = 1
	p.predictedMove <- idx
	return nil
}
