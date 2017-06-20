package balanced_brackets

import "testing"

func TestStackBalanced(t *testing.T) {

	got := Balance("{[()]}")

	want := true

	t.Error("Should be %q but is %q", want, got)

}

func TestStackUnbalanced(t *testing.T) {

	got := Balance("{(])}")

	want := false

	t.Error("Should be %q but is %q", want, got)
}
