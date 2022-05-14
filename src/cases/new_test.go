package cases

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	i := new([2]int)
	fmt.Println(i)
	j := [2]int{}
	fmt.Println(j)
	g := make([]int, 2)
	fmt.Println(g)

	x := new(string)
	fmt.Println(x)

	arr0 := [...]int{1, 2, 3} // array
	fmt.Println(arr0)
	arr1 := []int{1, 2, 3} // slice
	fmt.Println(arr1)
}
