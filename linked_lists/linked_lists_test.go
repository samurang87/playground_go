package linked_lists

import "testing"

func TestStack(t *testing.T) {

	a_node := &Node{data:"string a"}

	b_node := &Node{data:"string b"}

	linkedlist := LinkedList{}

	linkedlist.Stack(a_node)

	linkedlist.Stack(b_node)

	got := linkedlist.head.data

	want := "string b"

	if got != want {

		t.Error("Wanted %q, got %q", want, got)
	}
}