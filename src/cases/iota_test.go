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

func Test1(t *testing.T) {
	t.Log(AA, BB, CC, DD)
}
