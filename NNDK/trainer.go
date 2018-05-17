package nndk

import (
	G "github.com/chewxy/gorgonia"
)

type tictactoe struct{}

func (ttt *tictactoe) ReadInputVector(G *G.ExprGraph) (*G.Node, error) {
	return nil, nil
}
func (ttt *tictactoe) WriteComputedVector(n *G.Node) error {
	return nil
}
func (ttt *tictactoe) GetComputedVectors() G.Nodes {
	return nil
}
func (ttt *tictactoe) GetExpectedValue(offset int) (int, error) {
	return 0, nil
}

func (ttt *tictactoe) GetTrainer() (*tictactoe, error) {
	return ttt, nil
}
