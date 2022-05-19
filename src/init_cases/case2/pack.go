package main

import (
	"fmt"
)

var Pack = func() string {
	fmt.Println("get Pack")
	return "pack"
}()

func init() {
	fmt.Println("init pack start")
	a := Util
	fmt.Println("init pack done", a)
}
