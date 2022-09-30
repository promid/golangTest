package time_cases

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	t.Logf("now            : %v", time.Now())
	ti := <-timer.C
	t.Logf("timer expired at %v", ti)
	timer.Reset(2 * time.Second)
	ti = <-timer.C
	t.Logf("timer expired at %v", ti)
}
