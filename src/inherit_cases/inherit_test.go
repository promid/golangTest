package inherit_cases

import (
	"fmt"
	"testing"
)

func TestInherit(t *testing.T) {
	b := &b{}
	b.name = "world"
	b.hello()

	c := &c{}
	c.a.name = "world"
	c.a.hello()
	// wrong ↓
	// c.hello()

	d := &d{}
	d.a = &a{name: "world"}
	d.a.hello()
	// wrong ↓
	// d.hello()

}

type a struct {
	name string
}

func (a *a) hello() {
	fmt.Println("hello " + a.name)
}

type b struct {
	a
}

type c struct {
	a a
}

type d struct {
	a *a
}
