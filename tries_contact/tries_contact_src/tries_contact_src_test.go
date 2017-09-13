package tries_contact_src

import (
	"fmt"
	"testing"
)

func TestAddSingleWord(t *testing.T){

	var trie Node

	trie.AddEntry("bla")

	for _, child := range(trie.Children) {

		fmt.Println(child.Character)
	}
}