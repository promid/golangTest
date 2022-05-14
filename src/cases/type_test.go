package cases

import (
	"fmt"
	"testing"
)

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Printf("%d is an int value\n", arg)
		case int64:
			fmt.Printf("%d is an int64 value\n", arg)
		case string:
			fmt.Printf("%s is a string value\n", arg)
		default:
			fmt.Printf("%v is an unknown type\n", arg)
		}
	}
}

func TestMyPrintf(t *testing.T) {
	var v1 int = 1
	var v2 int64 = 64
	var v3 string = "hello"
	var v4 int32 = 2
	MyPrintf(v1, v2, v3, v4)
}
