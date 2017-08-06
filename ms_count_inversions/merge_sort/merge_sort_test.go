package merge_sort

import (
	"testing"
	"fmt"
)

func TestEven(t *testing.T) {

	i := []int{1, 6, 3, 4}

	r := []int{1, 3, 4, 6}


	if testEq(Sort(i), r) == false {

		t.Errorf("The result should be %v, got %v instead", r, i)
	}

}

func TestOdd(t *testing.T) {

	i := []int{1, 6, 3, 4, 2}

	r := []int{1, 2, 3, 4, 6}

	if testEq(Sort(i), r) == false {

		t.Errorf("The result should be %v, got %v instead", r, i)
	}


}

func TestOrderedSad(t *testing.T) {

	a := isOrdered([]int{1, 6, 3, 4})

	if a == true {

		t.Errorf("Is {1, 6, 3, 4} ordered? Should be false, got %v", a)
	}
}

func TestOrderedHappy(t *testing.T) {

	b := isOrdered([]int{1, 3, 4, 6})

	if b == false {

		t.Errorf("Is {1, 3, 4, 6} ordered? Should be true, got %v", b)
	}
}

func TestMerge(t *testing.T) {

	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}

	got := Merge(a, b)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}

	if testEq(got, want) == false {

		t.Errorf("Wanted %v, got %v", want, got)
	}

}

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		fmt.Println("One is nil and one is not!")
		return false
	}

	if len(a) != len(b) {
		fmt.Println("Oh noes, lengths differ!")
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			fmt.Println("Oh noes, %v and %v have different elements!", a, b)
			return false
		}
	}

	return true
}
