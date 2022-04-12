package main

import (
	"fmt"
)

func main() {
	doneCh := make(chan struct{}, 10)
	//close(doneCh)
	doneCh <- struct{}{}
	select {
	case <- doneCh:
		fmt.Println("1")
	default:
		fmt.Println("2")
	}
}
