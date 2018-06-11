package main

import (
	"context"
	"fmt"
)

func main() {
	for {
		model, err := getModel()
		if err != nil {
			panic(err)
		}
		fmt.Println("Let's play a game")
		//model := lstm.NewModel(9, 9, 100)
		p := NewPlayer()
		err = model.Predict(context.TODO(), p)
		if err != nil {
			panic(err)
		}
	}
}
