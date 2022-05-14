package cases

import (
	"fmt"
	"sync"
	"testing"
)

var lock sync.Mutex

func TestLock(t *testing.T) {
	chList := make([]chan struct{}, 10, 10)
	for i := 0; i < 10; i++ {
		i := i
		lock := lock
		chList[i] = make(chan struct{})
		go func(ch chan struct{}) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Printf("%d\n", i)
			ch <- struct{}{}
		}(chList[i])
	}
	for i := 0; i < 10; i++ {
		<-chList[i]
	}
}
