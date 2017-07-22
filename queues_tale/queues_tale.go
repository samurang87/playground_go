package main

import (
	"bufio"
	"os"
	"github.com/samurang87/playground_go/queues_tale/queues"
	"strings"
)

func Read() (input []string){

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		input = append(input, scanner.Text())
	}

	return input
}

func main() {

	input := Read()

	for _, item := range input[1:] {

		var q queues.Queue

		q.ExecuteFunction(strings.Split(item, " "))
		
	}


	}
