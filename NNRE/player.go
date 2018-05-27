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
		board: make([]int, 9),
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
			if move > 8 {
				continue
			}
			if p.board[move] == 1 {
				continue
			}
			p.offset++
			c <- move
			p.play = false
			c <- <-wait
		}
	}()
	return p
}

func (p *Player) Read() ([]float32, error) {
	move := <-p.c
	p.board[move] = 1
	output := make([]float32, 9)
	output[move] = 1
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
		if v > max && p.board[i] == 0 {
			max = v
			idx = i
		}
	}
	if idx == -1 {
		return errors.New("game end")
	}
	fmt.Println("My move:", idx)
	p.board[idx] = 1
	p.wait <- idx
	return nil
}
