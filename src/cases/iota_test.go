package cases

import (
	"testing"
)

const (
	AA = iota
	BB
	CC
	DD
)

const (
	mutexLocked      = 1 << iota // 1 左移 0
	mutexWoken                   // 1 左移 1 => 10 = 2
	mutexStarving                // 1 左移 2 => 100 = 4
	mutexWaiterShift = iota      // 3

	starvationThresholdNs = 1e6
)

func Test1(t *testing.T) {
	t.Log(AA, BB, CC, DD)
	t.Log(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)
	t.Log(starvationThresholdNs)
}
