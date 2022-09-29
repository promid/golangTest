package intro

import (
	"testing"
)

// case 1: 1^2 = 1
func TestXY1(t *testing.T) {
	x := 1
	y := 2
	expected := 1
	actual := exponent(x, y)
	if actual != expected {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}

// case
