package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
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
