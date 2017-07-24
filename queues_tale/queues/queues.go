package queues

import (
	"errors"
	"fmt"
	"bufio"
	"os"
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

func (q *Queue) enqueue(n *Node) {

	if q.isEmpty() {
		q.head = n
		q.tail = n

	} else {

		q.tail.next = n
		q.tail = n

	}

}

func (q *Queue) dequeue() (error){

	if q.isEmpty() {

		return errors.New("Cannot dequeue because empty")

	} else {

		q.head = q.head.next

		if q.head == nil {

			q.tail = nil

		}
	}

	return nil

}


func (q *Queue) printHead() {

	fmt.Println(q.head.data)

}

func (q *Queue) ExecuteFunction(l []string) {

	/*
	1 x: enqueue element  into the end of the queue.
	2: dequeue the element at the front of the queue.
	3: printHead the element at the front of the queue.
	*/

	switch length := len(l); length {

	case 2:

		v := l[1]

		n := new(Node)
		n.data = v

		q.enqueue(n)

	case 1:

		switch l[0] {

		case "2":

			q.dequeue()

		case "3":

			q.printHead()

		}

	}
}

func Read() (input []string){

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		input = append(input, scanner.Text())
	}

	return input
}

