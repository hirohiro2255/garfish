package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = add(slice)
	fmt.Println(slice)
}

func add(a []int) []int {
	a = append(a, 4)
	return a
}
