package quicksort_src

import (
	"fmt"
	"github.com/samurang87/playground_go"
)

func Partition(unsorted []int) (left []int, mid []int, right []int) {

	pivot := unsorted[0]

	for _, item := range unsorted {

		if item < pivot {

			left = append(left, item)

		} else if item > pivot {

			right = append(right, item)

		} else {

			mid = append(mid, item)

		}
	}

	return

}

func merge(collection [][]int) (merged []int) {

	for _, list := range collection {

		merged = append(merged, list...)
	}
	fmt.Println(hackerrank_utils.ListIntToStr(merged))
	return
}

func Quicksort(ar []int) (sorted []int) {

	if len(ar) <= 1 {

		return ar

	} else {

		left, mid, right := Partition(ar)

		return merge([][]int{Quicksort(left), mid, Quicksort(right)})

	}
}
