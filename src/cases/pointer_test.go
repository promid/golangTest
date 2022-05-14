package cases

import (
	"testing"
)

type Integer int

func (a *Integer) Add(b Integer) {
	*a += b
}

func TestAdd(t *testing.T) {
	var a Integer = 1
	a.Add(2)
	if a != 3 {
		t.Error("Add() failed, got", a, "expected 3")
	}
	t.Log("Add() passed")
}
