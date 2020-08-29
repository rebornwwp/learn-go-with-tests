package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreeting(t *testing.T) {
	buffer := bytes.Buffer{}

	Greeting(&buffer, "Chris")

	got := buffer.String()
	want := "hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
