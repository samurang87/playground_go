package queues

import (
	"errors"
	"fmt"
)

type Queue struct {
	head *Node
	tail *Node
}

type Node struct {
	data string
	next *Node
}

func (q Queue) isEmpty() bool {

	if q == (Queue{}) {
		return true

	} else {
		return false
	}
}

func (q *Queue) Enqueue(n *Node) {

	q.tail.next = n
	q.tail = n

}

func (q *Queue) Dequeue() (error){

	if q.isEmpty() {

		return errors.New("Cannot dequeue because empty")

	} else {

		q.head = q.head.next
	}

	return nil
}


func (q *Queue) Print() {

	fmt.Println(q.head.data)

}
