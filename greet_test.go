package greet_test

import (
	"bytes"
	"testing"

	"github.com/Zipeer2/greet"
)

func TestGreetUser_PromptsUserForANameAndRendersGreeting(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("Vlad")
	output := new(bytes.Buffer)
	greet.GreetUser(input, output)
	want := "Whats is your name?\nHello, Vlad\n"
	got := output.String()
	if want != got {
		t.Fatalf("wanted %q, got %q", want, got)
	}
}
