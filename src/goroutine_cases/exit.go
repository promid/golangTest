package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	child(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("main exits")
}

func child(ctx context.Context) {
	// grandchild will exit if child exits
	go grandchild(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child exist")
			return
		}
	}
}

func grandchild(ctx context.Context) {
	num := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("grandchild exists")
			return
		default:
			fmt.Println(num)
			num++
			time.Sleep(1 * time.Second)
		}
	}
}
