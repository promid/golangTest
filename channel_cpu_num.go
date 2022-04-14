package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum)
	c := make(chan bool, cpuNum)
	for i := 0; i < cpuNum; i++ {
		go Go(c, i)
	}

	for i := 0; i < cpuNum; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000; i++ {
		a += 1
	}
	fmt.Println(index, a)
	c <- true
}
