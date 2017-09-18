package tries_contact_src

<<<<<<< HEAD
import "strings"

type Trie struct {
	Children []*Node
}

type Node struct {
	Character string
	Children []*Node
	CompletesWord bool
}

func (trie *Node) AddEntry(word string){

	CharList := strings.Split(word, "")

	for _, Char := range CharList {

		// make sure you change CompletesWord accordingly!

		// if the character is not in the list of children

		// create a new Node in Children

		// do AddEntry for that Node


	}

}
=======
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
>>>>>>> fdacb60b3e002801d6e3082a51a143ed7811ad72
