package game

import (
	"log"

	"github.com/gonum/graph/path"
	"github.com/gonum/graph/simple"
)

type tttBoard [9]int

const (
	tokenX = 1
	tokenO = 0
)

// Generate all the positions
// 1 starts the game
func Generate() chan []int {
	c := make(chan []int, 0)
	go func() {
		for iter := 0; iter < 100; iter++ {
			log.Println("Iter", iter)
			var board tttBoard
			//for i := 0; i < 19683; i++ {
			for i := 0; i < 512; i++ {
				m := i
				for j := 0; j < 9; j++ {
					board[j] = m % 2
					m = m / 2
				}
				if isValid(board) && !isWinning(tokenX, board) {
					generateMoves(tokenX, board, c)
				}
			}
		}
		close(c)
	}()
	return c
}

func isValid(board tttBoard) bool {
	var tokenx, tokeno int
	for i := range board {
		switch board[i] {
		case tokenX:
			tokenx++
		case tokenO:
			tokeno++
		default:
		}
	}
	if tokenx == 5 && tokeno == 4 {
		return true
	}
	return false
}

func isWinning(token int, board tttBoard) bool {
	if (board[0] == token && board[0] == board[1] && board[1] == board[2]) ||
		(board[3] == token && board[3] == board[4] && board[4] == board[5]) ||
		(board[6] == token && board[6] == board[7] && board[7] == board[8]) ||
		(board[0] == token && board[0] == board[3] && board[3] == board[6]) ||
		(board[1] == token && board[1] == board[4] && board[4] == board[7]) ||
		(board[2] == token && board[2] == board[5] && board[5] == board[8]) ||
		(board[0] == token && board[0] == board[4] && board[4] == board[8]) ||
		(board[2] == token && board[2] == board[4] && board[4] == board[6]) {
		return true
	}
	return false
}

// games is the an array of the number of games and board
func generateMoves(token int, board tttBoard, c chan []int) {
	//log.Println(board)
	// Generate a graph for the board
	g := simple.NewDirectedGraph(0, 0)
	// First loop to create all the node (to avoid orphan nodes)
	var moves int
	for i := range board {
		g.AddNode(simple.Node(i))
		moves++
	}
	// Second loop to link the nodes
	for i := range board {
		var dst int
		switch board[i] {
		case tokenX:
			dst = tokenO
		case tokenO:
			dst = tokenX
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
							out := make([]int, len(path))
							for i, n := range path {
								out[i] = n.ID()
							}
							if out[0] == 0 && out[1] == 4 {
								c <- out
							}
						}
					}
				}
			}
		}
	}
}
