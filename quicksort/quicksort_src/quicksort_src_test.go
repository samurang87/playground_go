package quicksort_src

import (
	"testing"
	"github.com/samurang87/playground_go"
)

func TestRegularCase(t *testing.T) {

	input := []int{4, 5, 3, 7, 2}

	want := []int{3, 2, 4, 5, 7}

	got := Partition(input)

	if hackerrank_utils.TestEq(want, got) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}
