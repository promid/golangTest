package main

import "fmt"

var Util = func() string {
	fmt.Println("get Util")
	return "util"
}()

func init() {
	fmt.Println("init util start")
	b := Pack
	fmt.Println("init util done", b)
}
