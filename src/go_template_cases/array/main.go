package main

import (
	"os"
	"text/template"
)

func parseFiles() {
	pwd, _ := os.Getwd()
	t1, err := template.ParseFiles(pwd + "/src/go_template_cases/array/arr.txt")
	if err != nil {
		panic(err)
	}
	_ = t1.Execute(os.Stdout, "[\"90\",\"91\"]")
	println()
}

func main() {
	parseFiles()
}
