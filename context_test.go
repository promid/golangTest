package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	waitForJobDoneWithTimeout(ctx, 2*time.Second)
}

func waitForJobDoneWithTimeout(ctx context.Context, timeout time.Duration) {
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
		return
	case <-jobDone:
		fmt.Println("job done")
	}
	return
}
