package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/owulveryck/lstm"
)

func main() {
	// Read the weights
	b, err := ioutil.ReadFile("tictactoe.bin")
	if err != nil {
		log.Fatal("Cannot read binary file", err)
	}
	model := new(lstm.Model)
	err = model.UnmarshalBinary(b)
	if err != nil {
		log.Fatal("Cannot restore software 2.0", err)
	}
	p := new(Player)
	err = model.Predict(context.TODO(), p)
	if err != nil {
		log.Println(err)
	}
}
