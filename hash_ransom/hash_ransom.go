package main

import (
	"github.com/samurang87/playground_go"
	"github.com/samurang87/playground_go/hash_ransom/hash_ransom_src"
	"strings"
	"fmt"
)

func main() {

	input := hackerrank_utils.Read()

	magazine := hash_ransom.CreateMap(strings.Split(input[1], " "))

	note := hash_ransom.CreateMap(strings.Split(input[2], " "))

	if hash_ransom.CompareMaps(magazine, note) {

		fmt.Println("Yes")

	} else {

		fmt.Println("No")
	}

}