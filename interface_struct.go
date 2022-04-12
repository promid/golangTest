package main

type I interface {
	run()
}

type S struct {

}

var _ I = &S{}

func (s S) run() {

}
func (s S) run2() {

}

type S1 struct {
	s1i I
}

func genS1() {
	s1 := S1{}
	s := S{}
	s1.s1i = s
}