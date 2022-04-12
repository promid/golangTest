package main

import (
	"fmt"
)

func main() {
	a := closure(1)(2)
	fmt.Println(a)
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
