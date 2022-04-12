package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := arr[1:3]
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", slice)

	var s []int // not initialized
	s = append(s, 1)
	fmt.Println(s)

	var b []int // not initialized
	// panic
	// b[0] = 1
	fmt.Println(b)
}
