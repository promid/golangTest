package main

import (
	"context"
	"fmt"
	"time"
)

func CatchPanic() {
	fmt.Println("开始处理异常")
	// 获取异常信息
	if err := recover(); err != nil {
		//  输出异常信息
		fmt.Println("error:", err)
	}
	fmt.Println("结束异常处理")
}

func main() {

	defer CatchPanic()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go func() {
		for {
			select {
			case val, ok := <-ctx.Done():
				fmt.Println("val in goroutine:", val) // {}
				fmt.Println("ok in goroutine:", ok)   // false
				return
			default:
				fmt.Println("main is running")
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		// catching panic is required
		// if goroutine panics, main will be dead as well
		defer CatchPanic()
		panic("panic")
	}()

	val, ok := <-ctx.Done()
	// read from a closed channel
	fmt.Println("val in main:", val) // {}
	fmt.Println("ok in main:", ok)   // false

	time.Sleep(3 * time.Second)
}
