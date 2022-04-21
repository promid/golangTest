package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRecoverPanic(t *testing.T) {
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

type User struct {
	Id   int
	Name string
	Age  int
}

func TestReflect2(t *testing.T) {
	u := User{
		Id:   0,
		Name: "user",
		Age:  1,
	}
	SetUser(&u)
	fmt.Println(u)
}

func SetUser(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}
	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("new name")
	}
}
