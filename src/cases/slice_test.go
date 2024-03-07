package cases

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice := arr[1:3]
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", slice)

	var s []int // not initialized
	s = append(s, 1)
	fmt.Println(s)

	var b []int // not initialized
	// panic
	// b[0] = 1
	fmt.Println(b)
}

func PingPong(s []int) {
	// 如果append后超限，s的地址不变，但s中存储的数组指针变了，因为数组扩容后变成了新的数组
	s = append(s, 3)
}

func TestSlice2(t *testing.T) {
	s := make([]int, 0)
	fmt.Println(s) // []
	PingPong(s)
	fmt.Println(s) // []

	s = make([]int, 2)
	fmt.Println(s) // [0]
	PingPong(s)
	fmt.Println(s) // [0 3]
}

func TestSliceAppendAddr(t *testing.T) {
	s := make([]int, 5, 5)
	arr := [5]int{1, 2, 3, 4, 5}
	t.Logf("s addr = %p", &s)
	t.Logf("arr addr = %p", &arr)
	s = arr[:2]
	t.Logf("s addr = %p", &s)                      // address does not change
	t.Logf("arr pointer stored in s addr = %p", s) // address does not change
	s = append(s, 3)
	t.Logf("%v", s)
	t.Logf("s addr = %p", &s)                      // address does not change
	t.Logf("arr pointer stored in s addr = %p", s) // address does not change
	s = append(s, 4)
	t.Logf("%v", s)
	t.Logf("s addr = %p", &s)                      // address does not change
	t.Logf("arr pointer stored in s addr = %p", s) // address does not change
	s = append(s, 5)
	t.Logf("%v", s)
	t.Logf("s addr = %p", &s)                      // address does not change
	t.Logf("arr pointer stored in s addr = %p", s) // address does not change
	s = append(s, 6)
	t.Logf("%v", s)
	t.Logf("s addr = %p", &s)                      // address does not change
	t.Logf("arr pointer stored in s addr = %p", s) // the pointer to the array changes because a new array is created for larger capacity
}

func TestSliceCopy(t *testing.T) {
	s1 := make([]int, 5, 5)
	t.Log(s1)
	t.Logf("s1 addr = %p", &s1)
	t.Logf("arr addr = %p", s1)
	s2 := []int{1, 2}
	copy(s1, s2)
	t.Log(s1)
	t.Logf("s1 addr = %p", &s1) // address does not change
	t.Logf("arr addr = %p", s1) // address does not change
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	copy(s1, s3)
	t.Log(s1)
	t.Logf("s1 addr = %p", &s1) // address does not change
	t.Logf("arr addr = %p", s1) // address does not change
}

func TestSliceAssign(t *testing.T) {
	var a []string
	var b []string
	c := []string{"1"}
	a = c
	copy(b, c) // b没有容量, 所以拷贝不进去
	d := make([]string, 0)
	copy(d, c) // d没有容量, 所以拷贝不进去
	e := make([]string, 1)
	copy(e, c) // e有容量, 可以拷贝进去
	t.Logf("a = %v", a)
	t.Logf("b = %v", b)
	t.Logf("d = %v", d)
	t.Logf("e = %v", e)

	/*
	   slice_test.go:94: a = [1]
	   slice_test.go:95: b = []
	   slice_test.go:96: d = []
	   slice_test.go:97: e = [1]
	*/
}
