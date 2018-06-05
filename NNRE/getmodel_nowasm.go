// +build !wasm

package main

import (
	"encoding/gob"
	"os"

	"github.com/owulveryck/lstm"
)

func getModel() (*lstm.Model, error) {
	f, err := os.Open("tictactoe.bin")
	if err != nil {
		return nil, err
	}
	model := new(lstm.Model)
	dec := gob.NewDecoder(f)
	err = dec.Decode(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
