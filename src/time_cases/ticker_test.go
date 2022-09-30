package time_cases

import (
	"context"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			t.Log("tick")
		case <-ctx.Done():
			return
		}
	}
}
