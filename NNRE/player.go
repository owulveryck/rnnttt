package main

import (
	"fmt"
	"log"
	"strconv"
)

// Player ...
type Player struct {
	c chan int
}

// NewPlayer ...
func NewPlayer() *Player {
	c := make(chan int, 0)
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
		}
	}()
	return &Player{
		c: c,
	}
}

func (p *Player) Read() ([]float32, error) {
	move := <-p.c
	output := make([]float32, 9)
	output[move] = 1
	return output, nil
}

func (p *Player) Write(v []float32) error {
	log.Println(v)
	return nil
}
