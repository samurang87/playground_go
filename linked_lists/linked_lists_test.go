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

func TestRemoveHead(t *testing.T) {

	a_node := &Node{data:"string a"}

	b_node := &Node{data:"string b"}

	c_node := &Node{data:"string c"}

	linkedlist := LinkedList{}

	linkedlist.Stack(a_node)

	linkedlist.Stack(b_node)

	linkedlist.Stack(c_node)

	linkedlist.Remove(c_node)

	got := linkedlist.head.data

	want := "string b"

	if got != want {

		t.Error("Wanted %q, got %q", want, got)
	}
}

func TestRemoveMiddleNode(t *testing.T) {

	a_node := &Node{data:"string a"}

	b_node := &Node{data:"string b"}

	c_node := &Node{data:"string c"}

	linkedlist := LinkedList{}

	linkedlist.Stack(a_node)

	linkedlist.Stack(b_node)

	linkedlist.Stack(c_node)

	linkedlist.Remove(b_node)

	got := c_node.next.data

	want := "string a"

	if got != want {

		t.Error("Wanted %q, got %q", want, got)
	}

}

func TestReverse(t *testing.T){

	a_node := &Node{data:"string a"}

	b_node := &Node{data:"string b"}

	c_node := &Node{data:"string c"}

	linkedlist := LinkedList{}

	linkedlist.Stack(a_node)

	linkedlist.Stack(b_node)

	linkedlist.Stack(c_node)

	linkedlist.Reverse()

	got := linkedlist.head.data

	want := "string a"

	if got != want {

		t.Error("Wanted %q, got %q", want, got)
	}



}
