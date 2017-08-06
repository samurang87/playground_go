package merge_sort


func MergeSort(a []int) (b []int) {

	b = a

	return b

}

func isOrdered(a []int) bool {

	l := len(a) -1

	for i := 0; i < l; i++ {

		if a[i] > a[i+1] {

			return false

		} else {

			continue
		}
	}

	return true

}
