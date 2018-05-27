package main

import (
	"context"
	"encoding/gob"
	"io"
	"log"
	"os"

	"github.com/owulveryck/lstm"
	G "gorgonia.org/gorgonia"
)

func main() {
	model := lstm.NewModel(9, 9, 500)
	learnrate := 1e-3
	l2reg := 1e-3
	clipVal := float64(5)
	solver := G.NewRMSPropSolver(G.WithLearnRate(learnrate), G.WithL2Reg(l2reg), G.WithClip(clipVal))

	tset := newTictactoe()
	pause := make(chan struct{})
	infoChan, errc := model.Train(context.TODO(), tset, solver, pause)
	iter := 1
	for infos := range infoChan {
		if iter%100 == 0 {
			log.Printf("%v: %v", iter, infos)
		}
		if iter%500 == 0 {

			// save the software 2.0
			f, err := os.OpenFile("tictactoe.bin", os.O_RDWR|os.O_CREATE, 0755)
			if err != nil {
				log.Println(err)
			}
			enc := gob.NewEncoder(f)
			err = enc.Encode(model)
			if err != nil {
				log.Println(err)
			}
			if err := f.Close(); err != nil {
				log.Println(err)
			}
		}

		iter++
	}
	err := <-errc
	if err == io.EOF {
		close(pause)
		//return
	}
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	// save the software 2.0
	f, err := os.OpenFile("tictactoe.bin", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
	}
	enc := gob.NewEncoder(f)
	err = enc.Encode(model)
	if err != nil {
		log.Println(err)
	}
	if err := f.Close(); err != nil {
		log.Println(err)
	}
}
