package cases

import (
	"fmt"
	"testing"
)

type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestGenerics(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

func TestLess(t *testing.T) {
	b := Less(1, 1.2)
	if b != false {
		t.Errorf("expect false, got %v", b)
	}
}

func Less[K ~int | ~int32 | ~int64 | ~float32 | ~float64](a, b K) bool {
	return a < b
}
