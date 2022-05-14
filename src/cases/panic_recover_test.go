package cases

import (
	"fmt"
	"testing"
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

func TestRecoverPanic(t *testing.T) {
	//  处理异常的函数
	defer CatchPanic()
	exceptionFun()
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
