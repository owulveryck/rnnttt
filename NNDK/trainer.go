package main

import (
	"fmt"

	"github.com/owulveryck/lstm/datasetter"
	G "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

type game struct {
	currentBoard  []int
	computerMoves G.Nodes
	offset        int
}

func (gme *game) ReadInputVector(g *G.ExprGraph) (*G.Node, error) {
	oneHotMove := make([]int, 9)
	inputTensor := tensor.New(tensor.WithShape(9), tensor.WithBacking(oneHotMove))
	node := G.NewVector(g, tensor.Float32, G.WithName(fmt.Sprintf("input_%v", gme.offset)), G.WithShape(9), G.WithValue(inputTensor))
	return node, nil
}
func (gme *game) WriteComputedVector(n *G.Node) error {
	gme.computerMoves = append(gme.computerMoves, n)
	return nil
}
func (gme *game) GetComputedVectors() G.Nodes {
	return gme.computerMoves
}
func (gme *game) GetExpectedValue(offset int) (int, error) {
	return 0, nil
}

type tictactoe struct{}

func (ttt *tictactoe) GetTrainer() (datasetter.Trainer, error) {
	return &game{}, nil
}
