package main

import (
	"fmt"
)

func main() {
	//  处理异常的函数
	defer func() {
		fmt.Println("开始处理异常")
		// 获取异常信息
		if err := recover(); err != nil {
			//  输出异常信息
			fmt.Println("error:", err)
		}
		fmt.Println("结束异常处理")
	}()
	exceptionFun()
}

func exceptionFun() {
	fmt.Println("exceptionFun开始执行")
	panic("异常信息")
	fmt.Println("exceptionFun执行结束")
}
