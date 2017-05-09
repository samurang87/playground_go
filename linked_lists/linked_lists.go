package linked_lists

import "fmt"

//The new node is always added before the head of the given Linked List.
//And newly added node becomes the new head of the Linked List

type LinkedList struct {
	head *Node
}

type Node struct {
	data string
	next *Node
}

func (l *LinkedList) Stack(new_node *Node) {

	new_node.next = l.head
	l.head = new_node

}

func (l *LinkedList) Remove(to_be_removed *Node) {

	// if the current head is the one to be removed

	if l.head == to_be_removed {

		// point the head to its next

		l.head = to_be_removed.next

	} else {

		current_node := l.head

		var next_node *Node

		for {
			next_node = current_node.next

			if next_node == to_be_removed {

				break

			} else {

				current_node = next_node
			}
		}

		current_node.next = next_node.next
	}

}

func (l *LinkedList) GetLastTwo() (*Node, *Node) {

	current_node := l.head

	var next_node *Node

	for {
		next_node = current_node.next

		if next_node.next == nil {

			break
		}

	}

	return current_node, next_node

}


func (l *LinkedList) Reverse() {

	var previous_node *Node
	var next_node *Node
	current_node := l.head

	for {
		fmt.Println(current_node.data)
		next_node = current_node.next
		current_node.next = previous_node

		if next_node.next == nil {

			break

		} else {
			current_node = next_node

		}
	}

	l.head = next_node

}
