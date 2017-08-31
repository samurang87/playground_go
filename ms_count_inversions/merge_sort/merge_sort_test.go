package merge_sort

import (
	"testing"
	"github.com/samurang87/playground_go"
)

func TestEven(t *testing.T) {

	var mc MergeCounter

	a := []int{11, 1, 6, 3, 99, 123, 4}

	r := []int{1, 3, 4, 6, 11, 99, 123}


	if hackerrank_utils.TestEq(mc.Sort(a), r) == false {

		t.Errorf("The result should be %v, got %v instead", r, a)
	}

}

func TestOdd(t *testing.T) {

	var mc MergeCounter

	a := []int{1, 6, 3, 4, 2}

	r := []int{1, 2, 3, 4, 6}

	if hackerrank_utils.TestEq(mc.Sort(a), r) == false {

		t.Errorf("The result should be %v, got %v instead", r, a)
	}

}

func TestZeroInversions(t *testing.T) {

	got := CountInversions([]int{1, 1, 2, 2, 2})

	want := 0

	if want != got {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func TestSomeInversions(t *testing.T) {

	got := CountInversions([]int{2, 1, 3, 1, 2})

	want := 4

	if want != got {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func TestMerge(t *testing.T) {

	var mc MergeCounter

	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}

	got := mc.Merge(a, b)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if hackerrank_utils.TestEq(got, want) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}
