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

func TestPrint(t *testing.T) {

	var q Queue
	q.ExecuteFunction([]string{"1", "node_a"})
	q.ExecuteFunction([]string{"1", "node_b"})
	q.ExecuteFunction([]string{"2"})
	q.ExecuteFunction([]string{"3"})

}

func TestWithHackerrankInput(t *testing.T) {

	input := [][]string{
		{"1", "42"},
		{"2"},
		{"1", "14"},
		{"3"},
		{"1", "28"},
		{"3"},
		{"1", "60"},
		{"1", "78"},
		{"2"},
		{"2"},
	}

	var q Queue

	for _, list := range input {

		q.ExecuteFunction(list)

	}

}
