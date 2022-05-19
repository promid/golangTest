package main

import (
	"fmt"
)

// 初始化顺序：变量初始化->init()->main()

var T = a() // 1st

func init() { // 2nd
	fmt.Println("init in main.go ")
}

func a() int64 {
	fmt.Println("calling a()")
	return 2
}

func main() { // 3rd
	fmt.Println("calling main")
}
