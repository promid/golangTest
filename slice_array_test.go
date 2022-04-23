package main

import (
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
	arr := [3]int{1, 2, 3}
	slice := arr[:3]
	t.Log(slice)
	replaceSliceItem(slice) // correct
	t.Log(slice)
}

func replaceSliceItem(arr []int) {
	arr[0] = 4
}
