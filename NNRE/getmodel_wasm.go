// +build wasm

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net/http"
	"syscall/js"
	"time"

	"github.com/johanbrandhorst/fetch"
	"github.com/owulveryck/lstm"
	"github.com/vincent-petithory/dataurl"
)

func getModel() (*lstm.Model, error) {
	files := js.Global.Get("document").Call("getElementById", "knowledgeFile").Get("files")
	fmt.Println("file", files)
	fmt.Println("Length", files.Length())
	if files.Length() == 1 {
		fmt.Println("Reading from uploaded file")
		reader := js.Global.Get("FileReader").New()
		reader.Call("readAsDataURL", files.Index(0))
		for reader.Get("readyState").Int() != 2 {
			fmt.Println("Waiting for the file to be uploaded")
			time.Sleep(1 * time.Second)
		}
		content := reader.Get("result").String()
		dataURL, err := dataurl.DecodeString(content)
		if err != nil {
			return nil, err
		}
		model := new(lstm.Model)
		dec := gob.NewDecoder(bytes.NewReader(dataURL.Data))
		err = dec.Decode(model)
		if err != nil {
			fmt.Println("Error", err)
			return nil, err
		}

		return model, nil
	}
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
