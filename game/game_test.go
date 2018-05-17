package game

import "fmt"

func ExampleGenerate() {
	c := Generate()
	for v := range c {
		fmt.Println(v)
	}
}
