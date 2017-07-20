package queues

import "testing"

func TestEnqueue(t *testing.T) {

	var q Queue
	q.ExecuteFunction([]string{"1", "node_a"})
	q.ExecuteFunction([]string{"1", "node_b"})

	TestHead := q.head.data
	TestTail := q.tail.data

	if TestHead != "node_a" || TestTail != "node_b" {

		t.Errorf("Should be node_a and node_b, is %v and %v instead", TestHead, TestTail)
	}

}

func TestDequeue(t *testing.T) {

	var q Queue
	q.ExecuteFunction([]string{"1", "node_a"})
	q.ExecuteFunction([]string{"1", "node_b"})
	q.ExecuteFunction([]string{"2"})

	TestHead := q.head.data
	TestTail := q.tail.data

	if TestHead != "node_b" || TestTail != "node_b" {

		t.Errorf("Should be node_b and node_b, is %v and %v instead", TestHead, TestTail)
	}
}


