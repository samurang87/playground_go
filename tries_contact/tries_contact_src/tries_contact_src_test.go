package tries_contact_src

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
