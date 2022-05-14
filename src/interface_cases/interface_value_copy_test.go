package main

import (
	"testing"
)

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Connect() {
	println("Connecting to", pc.name)
}

func Test1(t *testing.T) {
	pc := PhoneConnector{name: "Phone"}
	// 将对象赋值给接口时会发生拷贝，而接口内部存储的是指向这个复制品的指针，而不是原对象的指针，所以既无法修改原对象的状态，也无法获取其指针
	c := Connector(pc)
	c.Connect()
	pc.name = "Phone2"
	c.Connect() // pc.name is still "Phone"
}

type Connector2 interface {
	Connect()
}

type PhoneConnector2 struct {
	name string
}

func (pc *PhoneConnector2) Connect() {
	println("Connecting to", pc.name)
}

func Test2(t *testing.T) {
	pc := &PhoneConnector2{name: "Phone"}
	c := Connector2(pc)
	c.Connect()
	pc.name = "Phone2"
	c.Connect() // pc.name is "Phone2"
}
