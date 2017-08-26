package quicksort_src

import (
	"testing"
	"github.com/samurang87/playground_go"
)

func TestPartitionRegularCase(t *testing.T) {

	input := []int{4, 5, 3, 7, 2}

	want := []int{3, 2, 4, 5, 7}

	got_left, got_mid, got_right := Partition(input)

	got := merge([][]int{got_left, got_mid, got_right})

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func TestQuicksortRegularCase(t *testing.T){

	input := []int{4, 5, 3, 7, 2}

	want := []int{2, 3, 4, 5, 7}

	got := Quicksort(input)

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}
}
