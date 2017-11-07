package main

import (
	"fmt"
	"github.com/samurang87/playground_go"
)

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	var array []int

	for i := 0; i < a; i++ {

		var c int

		fmt.Scan(&c)

		array = append(array, c)

	}

	for j := 0; j < b; j++ {

		array = append(array[1:], array[0])

	}

	fmt.Println(hackerrank_utils.ListIntToStr(array))
}