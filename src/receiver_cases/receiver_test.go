package receiver_cases

import (
	"fmt"
	"testing"
)

type coder interface {
	changeLanguage(lan string)
	code()
	language() string
}

// var _ coder = (*Gopher)(nil)

type Gopher struct {
	programmingLanguage string
}

func (g Gopher) changeLanguage(lan string) {
	g.programmingLanguage = lan
}

func (g Gopher) language() string {
	return g.programmingLanguage
}

func (g *Gopher) code() {
	fmt.Println("coding with " + g.programmingLanguage)
}

func TestChangeLanguage(t *testing.T) {
	// var g coder = Gopher{programmingLanguage: "go"}
	// ↑↑↑ Cannot use 'Gopher{programmingLanguage: "go"}' (type Gopher) as the type coder Type does not implement 'coder' as the 'code' method has a pointer receiver
	// g.changeLanguage("python")
	// g.code()

	var gPtr coder = &Gopher{programmingLanguage: "go"}
	gPtr.changeLanguage("python")
	t.Logf("gopher's language is %s", gPtr.language()) // go
	gPtr.code()
}
