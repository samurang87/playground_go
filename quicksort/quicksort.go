package main

import (
	"strings"
	"fmt"
	"github.com/samurang87/playground_go"
)

func main() {

	input := hackerrank_utils.Read()

	for _, item := range input[1:] {

		d := strings.Split(item, " ")
		fmt.Println(d)


	}

}
