package channel_cases

import (
	"fmt"
	"testing"
)

func TestChannel(t *testing.T) {
	doneCh := make(chan int, 2)
	defer close(doneCh)
	doneCh <- 1

	select {
	case <-doneCh:
		fmt.Println("1")
	default:
		fmt.Println("2")
	}
}

func TestParallel(t *testing.T) {
	length := 10
	chList := make([]chan struct{}, length, length)
	for i := 0; i < length; i++ {
		doneCh := make(chan struct{})
		chList[i] = doneCh
		go add(i, i, chList[i])
	}
	for i := 0; i < length; i++ {
		<-chList[i]
	}
}

func add(i, j int, doneCh chan struct{}) {
	fmt.Println(i + j)
	doneCh <- struct{}{}
}
