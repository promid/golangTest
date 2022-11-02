package go_template_cases

import (
	"bytes"
	"os"
	"testing"
	"text/template"
	"time"

	"k8s.io/apimachinery/pkg/util/rand"
)

func parseFiles() {
	t1, err := template.ParseFiles("./test.html")
	if err != nil {
		panic(err)
	}
	_ = t1.Execute(os.Stdout, "hello world")
	println()
}

type person struct {
	// Name must be exported
	Name string
}

func parseString() {
	str := "hello {{ .Name }}"
	p := person{
		Name: "bei",
	}
	t, _ := template.New("1").Parse(str)
	buffer := &bytes.Buffer{}
	_ = t.Execute(buffer, p)
	println(buffer.String())
}

func parseString2() {
	str := "{{ .placeholder1 }}, {{ .placeholder2 }}"
	vars := make(map[string]string)
	vars["placeholder1"] = "1011"
	vars["placeholder2"] = "1213"
	t, _ := template.New("2").Parse(str)
	buffer := &bytes.Buffer{}
	_ = t.Execute(buffer, vars)
	println(buffer.String())
}

func TestGoTemplate1(t *testing.T) {
	parseFiles()
	parseString()
	parseString2()
}

type Templ struct {
	Global GlobalTemplateValues
}

type GlobalTemplateValues struct {
	StartTimeStamp int64
	Rand           string
	Custom         map[string]string
}

func parseTemplateValues() {
	str := "{{ .Global.StartTimeStamp }}, {{ .Global.Rand }}, {{ .Global.Custom.A }}"
	t, _ := template.New("").Parse(str)
	buffer := &bytes.Buffer{}
	rand.Seed(time.Now().Unix())
	templateValues := Templ{
		Global: GlobalTemplateValues{
			StartTimeStamp: time.Now().Unix(),
			Rand:           rand.SafeEncodeString(rand.String(6)),
			Custom: map[string]string{
				"A": "a",
			},
		},
	}
	_ = t.Execute(buffer, templateValues)
	println(buffer.String())
}

func TestParseStructTemplateValues(t *testing.T) {
	parseTemplateValues()
}

func TestEscape(t *testing.T) {
	str := `{{ "{{ .Values.ABC }}" }}`
	vars := make(map[string]string)
	vars["placeholder1"] = "1011"
	tmpl, _ := template.New("").Parse(str)
	buffer := &bytes.Buffer{}
	_ = tmpl.Execute(buffer, vars)
	println(buffer.String())
}
