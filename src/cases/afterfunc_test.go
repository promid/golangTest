package cases

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("2...")
	})
	fmt.Println("1...")
	<-time.After(3 * time.Second)
	fmt.Println("3...")
}
