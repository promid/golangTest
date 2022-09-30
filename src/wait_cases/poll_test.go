package wait_cases

import (
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestPoll(t *testing.T) {
	fmt.Println("now :    " + time.Now().String())
	_ = wait.Poll(1*time.Second, 5*time.Second, func() (bool, error) {
		t.Log("hello at " + time.Now().String())
		time.Sleep(1 * time.Second)
		return false, nil
	})
	/*
		now :    2022-09-30 17:04:44.904554 +0800 CST m=+0.000705774
		hello at 2022-09-30 17:04:45.904851 +0800 CST m=+1.001241271
		hello at 2022-09-30 17:04:47.905203 +0800 CST m=+3.002021895
		hello at 2022-09-30 17:04:49.903854 +0800 CST m=+5.001042697
		hello at 2022-09-30 17:04:50.904263 +0800 CST m=+6.001617871
	*/
}

func TestPollImmediate(t *testing.T) {
	fmt.Println("now :    " + time.Now().String())
	_ = wait.PollImmediate(1*time.Second, 5*time.Second, func() (bool, error) {
		t.Log("hello at " + time.Now().String())
		time.Sleep(1 * time.Second)
		return false, nil
	})
	/*
		now :    2022-09-30 17:05:12.206666 +0800 CST m=+0.000767558
		hello at 2022-09-30 17:05:12.206928 +0800 CST m=+0.001028848
		hello at 2022-09-30 17:05:14.208046 +0800 CST m=+2.002185992
		hello at 2022-09-30 17:05:16.20715 +0800 CST m=+4.001318284
		hello at 2022-09-30 17:05:18.207984 +0800 CST m=+6.002171905
		hello at 2022-09-30 17:05:19.208068 +0800 CST m=+7.002262367
	*/
}
