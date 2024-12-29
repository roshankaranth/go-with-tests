package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	//Buffer type which is a struct, implements the Writer interface becasue it has the Write
	//method
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//in this case buffer is the mock, we are using to test if the right thing is being printed.
