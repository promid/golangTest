package channel_cases

import (
	"testing"
)

func TestClosedChannelWrite(t *testing.T) {
	ch := make(chan int)
	close(ch)
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	ch <- 1
}

func TestClosedChannelWithCacheRead(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	// if a channel is closed, it can still be read
	for i := 0; i < 5; i++ {
		val, ok := <-ch
		t.Logf("%d %v", val, ok) // 1 true 2 true 0 false 0 false 0 false
	}
}

func TestClosedChannelWithCacheRead2(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)
	// if a channel is closed, it can still be read
	for i := 0; i < 5; i++ {
		val, ok := <-ch
		t.Logf("%d %v", val, ok) // 1 true 0 false 0 false 0 false 0 false
	}
}

func TestClosedChannelWithoutCacheRead(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for i := 0; i < 5; i++ {
		val, ok := <-ch
		t.Logf("%d %v", val, ok) // 1 true 0 false 0 false 0 false 0 false
	}
}
