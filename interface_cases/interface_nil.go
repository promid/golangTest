package main

import (
	"fmt"
)

func main() {
	var i interface{}
	// 只有当接口存储的类型和对象都为nil时，接口才等于nil
	fmt.Println(i == nil) // false
	fmt.Println(i)        // <nil>

	var p *int
	i = p                 // 存储类型是一个指针
	fmt.Println(i == nil) // false
	fmt.Println(i)        // <nil>
}
