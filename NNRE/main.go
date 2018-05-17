package main

import (
	"context"
	"log"

	"github.com/owulveryck/lstm"
)

func main() {
	/*
		b, err := ioutil.ReadFile("tictactoe.bin")
		if err != nil {
			log.Fatal("Cannot read binary file", err)
		}
		model := new(lstm.Model)
		err = model.UnmarshalBinary(b)
		if err != nil {
			log.Fatal("Cannot restore software 2.0", err)
		}
	*/
	model := lstm.NewModel(9, 9, 100)
	p := NewPlayer()
	err := model.Predict(context.TODO(), p)
	if err != nil {
		log.Println(err)
	}
}
