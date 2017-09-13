package tries_contact_src

import "fmt"

type Node struct {
	Character string
	Children []*Node
	CountTraversed int
}

func (node *Node) AddEntry(word string){

	char := string(word[0])

	node.CountTraversed++

	isChild, newChild := isChild(char,node)

	if ! isChild {

		newChild.Character = char

		node.Children = append(node.Children, newChild)
	}

	if len(word) > 1 {

		newChild.AddEntry(word[1:])

	} else {

		var endChild *Node

		endChild.Character = "*"

		newChild.Children = append(newChild.Children, endChild)
	}

}

func isChild(char string, node *Node) (result bool, next *Node){

	for _, child := range node.Children{

		if child.Character == char {

			return true, child
		}

	}

	return false, nil
}

func (node *Node) FindPartial(word string) {

	char := string(word[0])

	isChild, newChild := isChild(char,node)

	if isChild {

		if len(word) > 1 {

			newChild.FindPartial(word[1:])

		} else {

			fmt.Println(newChild.CountTraversed)
		}


	}


}