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

func (s *Stack) Push(n *Node) {

	n.next = s.head
	s.head = n

}

func (s *Stack) Pop() (*Node, error) {

	if s.isEmpty() {

		return nil, errors.New("Cannot pop from empty stack")

	} else {

		to_remove := s.head
		s.head = s.head.next
		return to_remove, nil
	}
}

func isOpen(s string) (bool, error) {

	switch s {

	case "{","[","(":
		return true, nil

	case ")", "]", "}":
		return false, nil
	}

	return false, errors.New("Invalid character")
}

func isPair(first, second string) bool {

	return (first == "{" && second == "}") || (first == "[" && second == "]") || (first == "(" && second ==")")

}

func Balance(brackets string) bool {

	SetBracket := strings.SplitN(brackets, "", -1)

	var stack Stack

	for _, char := range SetBracket {

		open, err := isOpen(char)

		if err != nil {

			return false
		}

		if open {

			n := new(Node)
			n.data = char

			stack.Push(n)

		} else {

			removed, err := stack.Pop()

			if err != nil {

				return false
			}

			if isPair(removed.data, char) == false {

				return false
			}


		}

	}

	if stack.isEmpty() {

		return true

	} else {

		return false
	}
}
