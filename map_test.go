package main

import (
	"fmt"
	"testing"
)

func fill() (a_cool_map map[string]string) {
	// This fixes it: a_cool_map = make(map[string]string)
	a_cool_map["key"] = "value"
	return
}

func TestMap(t *testing.T) {
	a_cool_map := fill()
	fmt.Println(a_cool_map)
}

func TestReplaceMapItem(t *testing.T) {
	m := map[string]string{"a": "a"}
	t.Log(m)
	replaceMapItem(m)
	t.Log(m)
}

func replaceMapItem(m map[string]string) {
	m["a"] = "b"
}
