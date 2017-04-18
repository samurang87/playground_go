package linked_lists

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
