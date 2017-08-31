package main

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/samurang87/playground_go/ms_count_inversions/merge_sort"
	"github.com/samurang87/playground_go"
)

func main() {

	input := hackerrank_utils.Read()

	for i, item := range input[1:] {

		if i % 2 != 0 {

			d := strings.Split(item, " ")

			var n []int

			for _, item := range d {

				num, err := strconv.Atoi(item)

				if err != nil {

					continue
				}

				n = append(n, num)
			}

			fmt.Println(merge_sort.CountInversions(n))

		}

	}

}


