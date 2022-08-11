package cases

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeConversion(t *testing.T) {
	t1, err := time.Parse(time.RFC3339Nano, "2022-02-09T02:21:12.599508108Z")
	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		fmt.Println(t1)
	}
}

type A struct {
	Name string
}

func TestConversion(t *testing.T) {
	a := A{Name: "a"}
	convert(a)
}

func convert(obj interface{}) {
	a1, ok1 := obj.(A)
	if ok1 {
		fmt.Println(a1.Name)
	}
	a2, ok2 := obj.(*A) // wrong
	if ok2 {
		fmt.Println(a2.Name)
	}
}

func TestConvert(t *testing.T) {
	var a = 1
	b := convert1(a)
	t.Logf("b = " + b)

	var c = 1
	d := convert2(c) // panic
	t.Logf("d = " + d)
}

func convert1(obj interface{}) string {
	ret, ok := obj.(string)
	if !ok {
		fmt.Printf("ret type = %T\n", ret)
		fmt.Printf("ret = %s\n", ret)
	}
	return ret
}
func convert2(obj interface{}) string {
	ret := obj.(string) // panic
	return ret
}

func TestConvert3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("err = %v", err)
		}
	}()
	var a = 1
	b := convert3(a)
	t.Logf("b = " + b)
}

func convert3(obj interface{}) string {
	ret := obj.(string) // panic
	return ret
}

func TestConvert4(t *testing.T) {
	a := convert4(1)
	b := convert4(2)
	t.Logf("a = %d, b = %d", a, b)
}

func convert4(obj interface{}) int64 {
	ret := obj.(int64) // panic
	return ret
}
