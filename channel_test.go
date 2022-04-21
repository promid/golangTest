package main

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	doneCh := make(chan int, 2)
	//close(doneCh)
	doneCh <- 1

	select {
	case <-doneCh:
		fmt.Println("1")
	default:
		fmt.Println("2")
	}
}
