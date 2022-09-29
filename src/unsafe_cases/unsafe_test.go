package unsafe_cases

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafePointer(t *testing.T) {
	bytes := []byte{104, 101, 108, 108, 111}
	p := unsafe.Pointer(&bytes)
	str := *(*string)(p)
	fmt.Println(str)
}
