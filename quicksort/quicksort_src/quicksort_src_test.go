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

func TestPartitionSingleElement(t *testing.T)  {

	input := []int{1}

	want := []int{1}

	got_left, got_mid, got_right := Partition(input)

	got := merge([][]int{got_left, got_mid, got_right})

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func TestPartitionTwoElements(t *testing.T)  {

	input := []int{1, 2}

	want := []int{1, 2}

	got_left, got_mid, got_right := Partition(input)

	got := merge([][]int{got_left, got_mid, got_right})

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func TestQuicksortRegularCase(t *testing.T){

	input := []int{5, 8, 1, 3, 7, 9, 2}

	want := []int{1, 2, 3, 5, 7, 8, 9}

	got := Quicksort(input)

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestQuicksortDifficult(t *testing.T){

	input := []int{233, 3, 68, 243, 12, 83, 222, 22, 191, 12, 394, 12, 33, 411}

	got := Quicksort(input)

	want := []int{3, 12, 12, 12, 22, 33, 68, 83, 191, 222, 233, 243, 394, 411}

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}
}
