package cases

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	a := closure(1)(2)
	fmt.Println(a)
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
