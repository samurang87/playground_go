package main

import (
	"github.com/samurang87/playground_go/queues_tale/queues"
	"strings"
)

func main() {

	input := queues.Read()

	var q queues.Queue

	for _, item := range input[1:] {

		d := strings.Split(item, " ")

		q.ExecuteFunction(d)

	}

}
