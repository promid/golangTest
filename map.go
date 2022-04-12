package main

import (
	"fmt"
)

func fill() (a_cool_map map[string]string) {
	// This fixes it: a_cool_map = make(map[string]string)
	a_cool_map["key"] = "value"
	return
}
func main() {
	a_cool_map := fill()
	fmt.Println(a_cool_map)
}
