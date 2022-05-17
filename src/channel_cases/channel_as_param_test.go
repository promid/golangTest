package channel_cases

import (
	"fmt"
	"testing"
)

func TestParam(t *testing.T) {
	c := make(chan string)
	fmt.Printf("%p\n", &c)
	go channelAsParam(c)
	val := <-c
	fmt.Println(val)
}

func channelAsParam(c chan string) {
	fmt.Printf("%p\n", &c)
	c <- "Hello"
}

func TestParam2(t *testing.T) {
	c := make(chan string)
	fmt.Printf("%p\n", &c)
	go channelAsParam2(&c)
	val := <-c
	fmt.Println(val)
}

func channelAsParam2(c *chan string) {
	fmt.Printf("%p\n", c)
	*c <- "Hello"
}
