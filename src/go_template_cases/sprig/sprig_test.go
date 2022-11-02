package sprig

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/Masterminds/sprig"
)

type SprigTest struct {
	Name string
}

func TestSprig(t *testing.T) {
	s := &SprigTest{
		Name: "bei",
	}
	templ, _ := template.New("").Funcs(sprig.TxtFuncMap()).Parse("{{ .Name | upper }}")
	buffer := &bytes.Buffer{}
	_ = templ.Execute(buffer, s)
	println(buffer.String())
}
