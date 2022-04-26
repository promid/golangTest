package inner

import (
	"fmt"
	_ "unsafe"
)

//go:linkname hello testProject/golinkname/outer.Hello
func hello() {
	fmt.Println("hello world")
}
