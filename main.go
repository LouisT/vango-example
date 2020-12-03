package main // import "vercel-vango.vercel.app/vango-example"

import (
	"fmt"

	"vercel-vango.vercel.app/vango-example/helloworld"
)

var (
	// Target is exported for testing.
	Target string
)

// Print "Hello, vango!".
func main() {
	fmt.Println(helloworld.Message(Target))
}

// Set the target to "vango".
func init() {
	Target = "vango"
}