package main

import (
	"context"
	"encoding/gob"
	"log"
	"os"

	"github.com/owulveryck/lstm"
)

func main() {
	f, err := os.Open("tictactoe.bin")
	if err != nil {
		log.Fatal("Cannot read binary file", err)
	}
	model := new(lstm.Model)
	dec := gob.NewDecoder(f)
	err = dec.Decode(model)
	if err != nil {
		log.Fatal("Cannot restore software 2.0", err)
	}
	//model := lstm.NewModel(9, 9, 100)
	p := NewPlayer()
	err = model.Predict(context.TODO(), p)
	if err != nil {
		log.Println(err)
	}
}
