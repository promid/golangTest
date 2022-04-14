package main

import (
	"fmt"
)

func PingPong(s []int) {
	s = append(s, 3) // 注意，append后，s被改变不再指向原来的地址
}
func main() {
	s := make([]int, 0)
	fmt.Println(s) // []
	PingPong(s)
	fmt.Println(s) // []

	s = make([]int, 2)
	fmt.Println(s) // [0]
	PingPong(s)
	fmt.Println(s) // [0 3]
}
