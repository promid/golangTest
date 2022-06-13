package context_cases

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	ctx := context.TODO()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, ok) // ...(time) true
}

func TestDeadline2(t *testing.T) {
	ctx := context.TODO()
	deadline, ok := ctx.Deadline()
	fmt.Println(deadline, ok) // 0001-01-01 00:00:00 +0000 UTC false
}
