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

type Cluster struct {
	Name    string
	Address string
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

func TestSprig2(t *testing.T) {
	s := &Cluster{
		Name: "xxx999",
	}
	templ, _ := template.New("").Funcs(sprig.TxtFuncMap()).Parse(`{{ .Address | default ("https//api.system.svc.TODO.abc.io" | replace "TODO" (trimPrefix "xxx" .Name) ) }}`)
	buffer := &bytes.Buffer{}
	_ = templ.Execute(buffer, s)
	println(buffer.String())
}
