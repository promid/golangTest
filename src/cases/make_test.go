package cases

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	a := make([]int, 0, 5) // 向其中添加切片元素不会触发扩容, 因为这个切片的容量为5, 但是没有现在没有存储任何元素
	b := make([]int, 5, 5) // 向其中添加切片元素不会触发扩容, 因为这个切片的容量为5, 其中已经存储了5个0
	printSlice(a)
	printSlice(b)
	a = append(a, 1)
	b = append(b, 1)
	println("after append")
	printSlice(a)
	printSlice(b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
