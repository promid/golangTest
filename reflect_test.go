package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
}
