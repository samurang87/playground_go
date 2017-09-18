package main

import (

	"strings"
	"github.com/samurang87/playground_go"
	"github.com/samurang87/playground_go/tries_contact/tries_contact_src"
)

func main() {

	input := hackerrank_utils.Read()

	var trie tries_contact_src.Node

	for _, item := range input[1:] {

		instruction := strings.Split(item, " ")

		if instruction[0] == "add" {

			trie.AddEntry(instruction[1])

		} else {

			trie.FindPartial(instruction[1])
		}
	}


}
