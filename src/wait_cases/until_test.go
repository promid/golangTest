package wait_cases

import (
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestUntil(t *testing.T) {
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()
	fmt.Println("now :    " + time.Now().String())
	wait.Until(func() {
		fmt.Println("hello at " + time.Now().String())
		time.Sleep(1 * time.Second)
	}, 1*time.Second, stopCh)
	/*
		now :    2022-09-30 16:59:07.738135 +0800 CST m=+0.000843564
		hello at 2022-09-30 16:59:07.738621 +0800 CST m=+0.001329539
		hello at 2022-09-30 16:59:09.739607 +0800 CST m=+2.002261870
		hello at 2022-09-30 16:59:11.739856 +0800 CST m=+4.002458464
	*/
}

func TestNonSlidingUntil(t *testing.T) {
	stopCh := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(stopCh)
	}()
	fmt.Println("now :    " + time.Now().String())
	wait.NonSlidingUntil(func() {
		fmt.Println("hello at " + time.Now().String())
		time.Sleep(1 * time.Second)
	}, 1*time.Second, stopCh)
	/*
		now :    2022-09-30 16:59:33.249274 +0800 CST m=+0.001132888
		hello at 2022-09-30 16:59:33.257714 +0800 CST m=+0.009573200
		hello at 2022-09-30 16:59:34.25811 +0800 CST m=+1.009942302
		hello at 2022-09-30 16:59:35.259242 +0800 CST m=+2.011048003
		hello at 2022-09-30 16:59:36.260408 +0800 CST m=+3.012187260
		hello at 2022-09-30 16:59:37.260552 +0800 CST m=+4.012305505
	*/
}
