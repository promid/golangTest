package receiver_cases

import "testing"

type Case struct {
	Name string
}

func (c *Case) ChangeName(name string) {
	c.Name = name
}

func TestChangeName(t *testing.T) {
	c := Case{Name: "a"}
	c.ChangeName("b")
	print(c.Name)
}
