package extemplate_test

import (
	"html/template"
	"os"
	"strings"

	"github.com/utking/extemplate"
	"github.com/utking/extemplate/examples"
)

func ExampleExtemplate_ParseDir() {
	xt := extemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})
	_ = xt.ParseDir("examples", []string{".tmpl"})
	_ = xt.ExecuteTemplate(os.Stdout, "child.tmpl", nil)
	/* Output: Hello from child.tmpl
	Hello from partials/question.tmpl */
}

func ExampleExtemplate_ParseFS() {
	xt := extemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})
	_ = xt.ParseFS(examples.TemplateFiles, []string{".tmpl"})
	_ = xt.ExecuteTemplate(os.Stdout, "child.tmpl", nil)
	/* Output: Hello from child.tmpl
	Hello from partials/question.tmpl */
}
