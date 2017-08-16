package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"github.com/samurang87/playground_go/ms_count_inversions/merge_sort"
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

	for i, item := range input[1:] {

		if i % 2 != 0 {

			d := strings.Split(item, " ")
			fmt.Println(d)

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


