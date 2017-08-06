package merge_sort

import "testing"

func TestEven(t *testing.T) {

	i := []int{1, 6, 3, 4}

	r := []int{1, 3, 4, 6}

	if testEq(i, r) == false {

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



func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}