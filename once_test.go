package main

import (
	"sync"
	"testing"
)

var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

func doPrint() {
	// once 的 Do()方法可以保证在全局范围内只调用指定的函数一次（这里指 setup() 函数）， 而且所有其他goroutine在调用到此语句时， 将会先被阻塞， 直至全局唯一的 once.Do()调用结束后才继续
	once.Do(setup)
	print(a)
}

func twoPrint() {
	go doPrint()
	go doPrint()
}

func TestOnce(t *testing.T) {
	twoPrint()
}
