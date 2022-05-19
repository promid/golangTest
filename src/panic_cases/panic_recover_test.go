package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRecoverPanic(t *testing.T) {
	//  处理异常的函数
	defer CatchPanic()
	exceptionFun()
}

func TestRecoverPanic2(t *testing.T) {
	//  处理异常的函数
	defer CatchPanic()
	go exceptionFun() // panic won't be caught in another goroutine
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		num := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
			default:
				fmt.Println(num)
				num++
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	cancel()
}

func exceptionFun() {
	fmt.Println("exceptionFun开始执行")
	panic("异常信息")
	fmt.Println("exceptionFun执行结束") // unreachable
}

func TestRecoverDeepInsidePanic(t *testing.T) {
	defer CatchPanic()
	a()
}

func a() {
	fmt.Println("a pre")
	bHasPanic()
	fmt.Println("a post") // unreachable
}

func bHasPanic() {
	panic("b has panic")
}
