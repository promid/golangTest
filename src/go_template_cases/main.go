package main

import (
	"bytes"
	"os"
	"text/template"
)

func parseFiles() {
	pwd, _ := os.Getwd()
	t1, err := template.ParseFiles(pwd + "/src/go_template_cases/test.html")
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

func main() {
	parseFiles()
	println()
	parseString()
	println()
	parseString2()
}
