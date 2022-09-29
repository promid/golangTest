package context_cases

import (
	"context"
	"fmt"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	if err := waitForJobDoneWithTimeout(ctx, 2*time.Second); err != nil {
		t.Error(err)
	}
}

func waitForJobDoneWithTimeout(ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	jobDone := make(chan struct{})
	go func() {
		fmt.Println("job started")
		time.Sleep(3 * time.Second)
		jobDone <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("timeout, exit")
	case <-jobDone:
		fmt.Println("job done")
	}
	return ctx.Err()
}

func TestTick(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	go tick(ctx)
	time.Sleep(5 * time.Second)
	cancel()
}

func tick(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-ctx.Done():
			return
		}
	}
}

func TestUntilWithContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	fmt.Println("start")
	wait.UntilWithContext(ctx, func(ctx context.Context) {
		fmt.Println("tick")
	}, 5*time.Second)
}

func TestContextPointer(t *testing.T) {
	ctx1 := context.TODO()
	ctx2 := context.TODO()
	if ctx1 != ctx2 {
		t.Fatalf("ctx1 != ctx2")
	}
}
