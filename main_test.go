package main

import (
	"testing"

	"vercel-vango.vercel.app/vango-example/helloworld"
)

func TestHelloWorld(t *testing.T) {
	if helloworld.Message("vango") != "Hello, vango!" {
		t.Fatal("Test fail")
	}
}

func TestTarget(t *testing.T) {
	if Target != "vango" {
		t.Fatal("Test fail")
	}
}