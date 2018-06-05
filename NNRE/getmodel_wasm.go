// +build wasm

package main

import (
	"encoding/gob"
	"net/http"

	"github.com/johanbrandhorst/fetch"
	"github.com/owulveryck/lstm"
)

func getModel() (*lstm.Model, error) {
	c := http.Client{
		Transport: &fetch.Transport{},
	}
	resp, err := c.Get("/tictactoe.bin")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	model := new(lstm.Model)
	dec := gob.NewDecoder(resp.Body)
	err = dec.Decode(model)
	if err != nil {
		return nil, err
	}

	return model, nil
}
