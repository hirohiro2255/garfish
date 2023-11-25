package main

import "fmt"

func main() {
	b, err := boardFromFen(DEFAULT_POS)
	if err != nil {
		fmt.Println(err)
	} else {
		b.printBoard()
	}
}
