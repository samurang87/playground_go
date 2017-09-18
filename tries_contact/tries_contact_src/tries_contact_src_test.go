package tries_contact_src

import (
	"testing"
)

func TestCount(t *testing.T){

	var trie Node

	trie.AddEntry("bla")

	trie.AddEntry("blub")

	trie.AddEntry("bar")

	got := trie.CountTraversed

	want := 3

	if want != got {

		t.Errorf("Wanted %v, got %v", want, got)
	}


}

func TestFind(t *testing.T){

	var trie Node

	trie.AddEntry("bla")

	trie.AddEntry("blub")

	trie.AddEntry("bar")

	trie.FindPartial("bl")

}
