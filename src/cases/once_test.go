package cases

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var once sync.Once
var once2 sync.Once

func doPrint(doneCh chan struct{}) {
	// once的Do()方法可以保证在全局范围内只调用指定的函数一次,而且所有其他goroutine在调用到此语句时,将会先被阻塞,直至全局唯一的once.Do()调用结束后才继续
	once.Do(func() {
		fmt.Println("hello world")
	})
	fmt.Println("after once.Do")
	doneCh <- struct{}{}
}

func twoPrint(doneCh chan struct{}) {
	go doPrint(doneCh)
	go doPrint(doneCh)
}

func TestOnce(t *testing.T) {
	doneCh := make(chan struct{}, 2)
	twoPrint(doneCh)
	for i := 0; i < 2; i++ {
		<-doneCh
	}
}

func doOnce2() {
	once2.Do(func() {
		fmt.Println("1")
		time.Sleep(2 * time.Second)
		fmt.Println("2")
	})
	// will wait until once2.Do() is done
	fmt.Println("3")
}

func TestOnce2(t *testing.T) {
	doOnce2()
}
