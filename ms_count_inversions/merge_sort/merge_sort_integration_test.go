package merge_sort

import (
	"strings"
	"strconv"
	"fmt"
	"testing"
	"github.com/samurang87/playground_go"
)

func TestIntegration(t *testing.T) {

	input := hackerrank_utils.ReadFromStdinLikeFile("../example_input_2.txt")

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

			fmt.Println(CountInversions(n))

		}

	}

}

