package main

type Connector2 interface {
	Connect()
}

type PhoneConnector2 struct {
	name string
}

func (pc *PhoneConnector2) Connect() {
	println("Connecting to", pc.name)
}

func main() {
	pc := &PhoneConnector2{name: "Phone"}
	c := Connector2(pc)
	c.Connect()
	pc.name = "Phone2"
	c.Connect() // pc.name is "Phone2"
}
