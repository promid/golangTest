package same_signature_interface

import (
	"fmt"
	"testing"
)

func TestChild(t *testing.T) {
	c := NewChild("xiaoming", 10, "male")
	fmt.Println("child name = ", c.Name())
	p := c.(Parent)
	fmt.Println("parent name = ", p.Name())
}

func TestParent(t *testing.T) {
	p := NewParent("qidong", 30)
	fmt.Println("parent name = ", p.Name())
	// wrong
	// panic: interface conversion: *sub_interface.Parent1 is not sub_interface.child: missing method Sex
	c := p.(Child)
	fmt.Println("child name = ", c.Name())
}
