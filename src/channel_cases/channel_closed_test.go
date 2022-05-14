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
	for i := 0; i < 5; i++ {
		t.Logf("%d ", <-ch) // 1 2 0 0 0
	}
}

func TestClosedChannelWithoutCacheRead(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for i := 0; i < 5; i++ {
		t.Logf("%d\n", <-ch)
	}
}
