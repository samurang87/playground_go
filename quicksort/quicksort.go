package main

import (
	"strings"
	"fmt"
	"github.com/samurang87/playground_go"
	"github.com/samurang87/playground_go/quicksort/quicksort_src"
	"strconv"
)

func main() {

	input := hackerrank_utils.Read()

	for _, item := range input[1:] {

		d := strings.Split(item, " ")

		var n []int

		for _, item := range d {

			num, err := strconv.Atoi(item)

			if err != nil {

				continue
			}

			n = append(n, num)
		}

		fmt.Println(hackerrank_utils.ListIntToStr(quicksort_src.Partition(n)))

	}

}
