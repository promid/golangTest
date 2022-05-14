package main

import (
	"fmt"
	"testing"
)

type I interface {
	run()
}

type S struct {
	name string
}

var _ I = &S{}

func (s S) run() {

}
func (s S) run2() {

}

type S1 struct {
	prop I
}

func TestInterfaceStruct(t *testing.T) {
	s1 := S1{}
	s := S{"s"}
	s1.prop = s

	switch v := s1.prop.(type) {
	case S:
		fmt.Println(v.name)
	}
}
