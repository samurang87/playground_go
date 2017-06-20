package balanced_brackets

import (
	"errors"
	"strings"
)



type Stack struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func (s Stack) isEmpty() bool {

	if s == (Stack{}) {

		return true

	} else {

		return false
	}
}

func (s *Stack) StackOn(n *Node) {

	n.next = s.head
	s.head = n

}

func (s *Stack) StackOff() (*Node, error) {

	if s.isEmpty() {

		return nil, errors.New("Cannot pop from empty stack")

	} else {

		to_remove := s.head
		s.head = s.head.next
		return to_remove, nil
	}
}

func IsOpen(s string) (bool, error) {

	switch s {

	case "{","[","(":
		return true, nil

	case ")", "]", "}":
		return false, nil
	}

	return nil, errors.New("Invalid character")
}

func Balance(brackets string) bool {

	SetBracket := strings.SplitN(brackets, "", -1)

	var stack Stack

	for _, char := range SetBracket {

		if IsOpen(char) {

			var n *Node
			n.data = char

			stack.StackOn(n)

		} else {



		}
	}



}
