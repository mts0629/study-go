package main

import (
	"os"
	"text/template"
)

func main() {
	// Text template
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n") // Action "{{.}}" will be replaced
	if err != nil {
		panic(err)
	}

	// template.Must panics on error of Parse
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Generate texts
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Helpler function
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// Template by field name
	t2 := Create("t2", "Name: {{.Name}}\n")

	// Replace by .Name field of struct
	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// Replace by "Name" key of map
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Micky Mouse",
	})

	// Conditional execution
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "") // Default value is considered as false

	// Loop for slices, arrays, maps or channels
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
