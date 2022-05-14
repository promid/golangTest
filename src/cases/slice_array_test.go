package cases

import (
	"fmt"
	"testing"
)

func TestReplaceArrayItem(t *testing.T) {
	// the length of an array  is fixed, can't be changed.
	arr := [3]int{1, 2, 3}
	t.Log(arr)
	replaceArrayItem(arr) //wrong
	t.Log(arr)
	replaceArrayItemCorrect(&arr) //correct
	t.Log(arr)
}

func replaceArrayItem(arr [3]int) {
	// 值传递,拷贝
	arr[0] = 4
}

func replaceArrayItemCorrect(arr *[3]int) {
	arr[0] = 4
}

func TestReplaceSliceItem(t *testing.T) {
	// [3]int 和 [4]int 是完全不同的类型; 数组不需要显式的初始化；数组的零值是可以直接使用的，数组元素会自动初始化为其对应类型的零值
	// arr := [...]int{1, 2, 3}
	arr := [3]int{1, 2, 3}
	// 一个切片是一个数组片段的描述。它包含了指向数组的指针，片段的长度， 和容量（片段的最大长度）。
	// 切片操作并不会复制底层的数组。整个数组将被保存在内存中，直到它不再被引用
	slice := arr[:3]
	t.Log(slice)
	t.Logf("slice addr = %p\n", &slice)
	replaceSliceItem(slice) // correct
	t.Log(slice)
	replaceSliceItem2(&slice) // correct
	t.Log(slice)
}

func replaceSliceItem(arr []int) {
	fmt.Printf("arr addr = %p\n", &arr)
	// 其实它也是个值传递，只不过它是从原数组中又切了一个一样的切片，修改它会改变原数组，自然也改变了从数组中切出来的切片
	// https://blog.go-zh.org/go-slices-usage-and-internals
	// https://studygolang.com/articles/9876
	arr[0] = 4
}

func replaceSliceItem2(arr *[]int) {
	fmt.Printf("arr addr = %p\n", arr)
	(*arr)[0] = 5
}
