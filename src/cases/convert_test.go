package cases

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeConversion(t *testing.T) {
	t1, err := time.Parse(time.RFC3339Nano, "2022-02-09T02:21:12.599508108Z")
	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		fmt.Println(t1)
	}
}

type A struct {
	Name string
}

func TestConversion(t *testing.T) {
	a := A{Name: "a"}
	convert(a)
}

func convert(obj interface{}) {
	a1, ok1 := obj.(A)
	if ok1 {
		fmt.Println(a1.Name)
	}
	a2, ok2 := obj.(*A) // wrong
	if ok2 {
		fmt.Println(a2.Name)
	}
}
