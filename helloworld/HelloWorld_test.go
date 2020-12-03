package helloworld

import (
	"testing"
)

func TestMessage(t *testing.T) {
	if Message("World") != "Hello, World!" {
		t.Fatal("Test fail")
	}
}