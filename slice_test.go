package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
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

func PingPong(s []int) {
	s = append(s, 3) // 注意，append后，s被改变不再指向原来的地址
}

func TestSlice2(t *testing.T) {
	s := make([]int, 0)
	fmt.Println(s) // []
	PingPong(s)
	fmt.Println(s) // []

	s = make([]int, 2)
	fmt.Println(s) // [0]
	PingPong(s)
	fmt.Println(s) // [0 3]
}
