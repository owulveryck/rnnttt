package main

import (
	"fmt"
	"log"

	"github.com/gonum/graph/path"
	"github.com/gonum/graph/simple"
)

type tttBoard [9]int

func main() {
	var board tttBoard
	for i := 0; i < 19683; i++ {
		c := i
		for j := 0; j < 9; j++ {
			board[j] = c % 3
			c = c / 3
		}
		if isValid(board) && isWinning(2, board) && !isWinning(1, board) {
			//			fmt.Println(board)
			generateMoves(2, board)
		}
	}
}

func isValid(board tttBoard) bool {
	var ones, twos int
	for i := range board {
		switch board[i] {
		case 2:
			twos++
		case 1:
			ones++
		default:
		}
	}
	diff := (twos - ones) * (twos - ones)
	if diff == 0 || diff == 1 {
		return true
	}
	return false
}

func isWinning(value int, board tttBoard) bool {
	if (board[0] == value && board[0] == board[1] && board[1] == board[2]) ||
		(board[3] == value && board[3] == board[4] && board[4] == board[5]) ||
		(board[6] == value && board[6] == board[7] && board[7] == board[8]) ||
		(board[0] == value && board[0] == board[3] && board[3] == board[6]) ||
		(board[1] == value && board[1] == board[4] && board[4] == board[7]) ||
		(board[2] == value && board[2] == board[5] && board[5] == board[8]) ||
		(board[0] == value && board[0] == board[4] && board[4] == board[8]) ||
		(board[2] == value && board[2] == board[4] && board[4] == board[6]) {
		return true
	}
	return false
}

// games is the an array of the number of games and board
func generateMoves(token int, board tttBoard) {
	// Generate a graph for the board
	g := simple.NewDirectedGraph(0, 0)
	// First loop to create all the node (to avoid orphan nodes)
	var moves int
	for i := range board {
		if board[i] != 0 {
			g.AddNode(simple.Node(i))
			moves++
		}
	}
	// Second loop to link the nodes
	for i := range board {
		var dst int
		switch board[i] {
		case 1:
			dst = 2
		case 2:
			dst = 1
		}
		for _, n := range g.Nodes() {
			if board[n.ID()] == dst {
				g.SetEdge(simple.Edge{
					F: simple.Node(i),
					T: n,
					W: 0,
				})
			}
		}
	}
	allPaths, ok := path.FloydWarshall(g)
	if ok != true {
		log.Fatal("Non ok")
	}
	for i := range board {
		if board[i] == token {
			for j := range board {
				if j != i && board[j] == token {
					paths, _ := allPaths.AllBetween(simple.Node(i), simple.Node(j))
					for _, path := range paths {
						if len(path) == moves {
							fmt.Println(path)
						}

					}
				}
			}
		}
	}

}

func isEqual(a []int, b tttBoard) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isNull(board tttBoard) bool {
	var total int
	for i := range board {
		total += board[i]
	}
	if total == 0 {
		return true
	}
	return false
}

func getNumberOf(token int, board tttBoard) int {
	var ret int
	for i := range board {
		if board[i] == token {
			ret++
		}
	}
	return ret
}