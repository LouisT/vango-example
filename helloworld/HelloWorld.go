package helloworld // import "vercel-vango.vercel.app/vango-example/helloworld"

import (
	"bytes"
	"html/template"
	"log"
)

// Message returns a formatted "Hello, World!" string.
func Message(name string) string {
	// Create a template.
	t := template.Must(template.New("HelloWorld").Parse("Hello, {{.Name}}!"))

	// Create a result buffer.
	var result bytes.Buffer

	// Execute the template, passing name data.
	if err := t.Execute(&result, struct{ Name string }{ Name: name }); err != nil {
		log.Fatal(err)
	}

	// Return a string.
	return result.String()
}