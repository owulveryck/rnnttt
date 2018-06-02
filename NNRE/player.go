package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// Player ...
type Player struct {
	board  []int
	c      chan int
	wait   chan int
	offset int
	play   bool
}

// NewPlayer ...
func NewPlayer() *Player {
	c := make(chan int, 0)
	wait := make(chan int, 0)
	p := &Player{
		c:     c,
		wait:  wait,
		board: make([]int, 18),
	}

	go func() {
		for {
			p.play = true
			fmt.Print("Enter move: ")
			var input string
			fmt.Scanln(&input)
			move, err := strconv.Atoi(input)
			if err != nil {
				log.Println(err)
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
			c <- move
			p.play = false
			c <- <-wait
		}
	}()
	return p
}

func (p *Player) Read() ([]float32, error) {
	<-p.c
	output := make([]float32, 18)
	for i := range p.board {
		output[i] = float32(p.board[i])
	}
	p.offset++
	return output, nil
}

func (p *Player) Write(v []float32) error {
	if p.play {
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
	p.board[idx+9] = 1
	p.wait <- idx
	return nil
}
