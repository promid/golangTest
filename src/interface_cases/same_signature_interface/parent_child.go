package same_signature_interface

type Parent interface {
	Name() string
	Age() int
}

type Child interface {
	Name() string
	Age() int
	Sex() string
}

type Child1 struct {
	name string
	age  int
	sex  string
}

func NewChild(name string, age int, sex string) Child {
	return &Child1{name: name, age: age, sex: sex}
}

func (c *Child1) Name() string {
	return c.name
}

func (c *Child1) Age() int {
	return c.age
}

func (c *Child1) Sex() string {
	return c.sex
}

var _ Child = &Child1{}

type Parent1 struct {
	name string
	age  int
}

func NewParent(name string, age int) Parent {
	return &Parent1{name: name, age: age}
}

func (p *Parent1) Name() string {
	return p.name
}

func (p *Parent1) Age() int {
	return p.age
}

var _ Parent = &Parent1{}
