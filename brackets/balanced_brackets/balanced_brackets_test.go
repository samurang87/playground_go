package balanced_brackets

import "testing"

func TestStackBalanced(t *testing.T) {

	got := Balance("{[]}()")

	want := true

	if got != want {

		t.Errorf("Should be %v but is %v", want, got)

	}

}

func TestStackUnbalanced(t *testing.T) {

	got := Balance("{(])}")

	want := false

	if got != want {

		t.Errorf("Should be %v but is %v", want, got)

	}
}

func TestOpenStack(t *testing.T) {

	got := Balance("(")

	want := false

	if got != want {

		t.Errorf("Should be %v but is %v", want, got)

	}
}
