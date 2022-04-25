package main

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func doPrint(doneCh chan struct{}) {
	// once的Do()方法可以保证在全局范围内只调用指定的函数一次,而且所有其他goroutine在调用到此语句时,将会先被阻塞,直至全局唯一的once.Do()调用结束后才继续
	once.Do(func() {
		fmt.Println("hello world")
	})
	fmt.Println("after once.Do")
	doneCh <- struct{}{}
}

func twoPrint(doneCh chan struct{}) {
	go doPrint(doneCh)
	go doPrint(doneCh)
}

func TestOnce(t *testing.T) {
	doneCh := make(chan struct{}, 2)
	twoPrint(doneCh)
	for i := 0; i < 2; i++ {
		<-doneCh
	}
}
