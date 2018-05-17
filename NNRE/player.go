package main

import (
	"fmt"
	"log"
	"strconv"
)

// Player ...
type Player struct {
	c    chan int
	wait chan struct{}
}

// NewPlayer ...
func NewPlayer() *Player {
	c := make(chan int, 0)
	wait := make(chan struct{}, 0)
	go func() {
		for {
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
			c <- move
			<-wait
		}
	}()
	return &Player{
		c:    c,
		wait: wait,
	}
}

func (p *Player) Read() ([]float32, error) {
	move := <-p.c
	output := make([]float32, 9)
	output[move] = 1
	return output, nil
}

func (p *Player) Write(v []float32) error {
	// Get the max probability
	max := float32(0)
	idx := -1
	for i, v := range v {
		if v > max {
			max = v
			idx = i
		}
	}
	fmt.Println("My move:", idx)
	p.wait <- struct{}{}
	return nil
}
