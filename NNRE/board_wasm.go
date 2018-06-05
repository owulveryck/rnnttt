// +build wasm

package main

import "fmt"

type board [9]string

func (b board) draw() {
	fmt.Println(b)
}
