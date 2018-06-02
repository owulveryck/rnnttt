package main

import (
	"fmt"
	"io"

	"github.com/owulveryck/lstm/datasetter"
	"github.com/owulveryck/rnnttt/game"
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type party struct {
	targetBoard   []int
	currentBoard  []int
	computerMoves G.Nodes
	offset        int
}

func (p *party) ReadInputVector(g *G.ExprGraph) (*G.Node, error) {
	if p.offset == len(p.targetBoard)-1 {
		return nil, io.EOF
	}
	oneHotMove := make([]float32, 18)
	for i, v := range p.targetBoard {
		if i <= p.offset {
			if v != 9 {
				oneHotMove[v+9*(i%2)] = float32(1)
			}
		}
	}
	//log.Printf("%v %v %v %v", p.targetBoard, p.offset, p.targetBoard[p.offset], oneHotMove)
	inputTensor := tensor.New(tensor.WithShape(18), tensor.WithBacking(oneHotMove))
	node := G.NewVector(g, tensor.Float32, G.WithName(fmt.Sprintf("input_%v", p.offset)), G.WithShape(18), G.WithValue(inputTensor))
	p.offset++
	return node, nil
}
func (p *party) WriteComputedVector(n *G.Node) error {
	p.computerMoves = append(p.computerMoves, n)
	return nil
}
func (p *party) GetComputedVectors() G.Nodes {
	return p.computerMoves
}
func (p *party) GetExpectedValue(offset int) (int, error) {
	return p.targetBoard[offset+1], nil
}

type tictactoe struct {
	c chan []int
}

func newTictactoe() *tictactoe {
	return &tictactoe{
		c: game.Generate(),
	}
}

func (ttt *tictactoe) GetTrainer() (datasetter.Trainer, error) {
	// movesChan is a channel fed with
	targetBoard := <-ttt.c
	return &party{
		targetBoard: append([]int{9}, targetBoard...),
	}, nil
}
