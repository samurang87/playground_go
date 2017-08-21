package quicksort_src

func Partition(unsorted []int) (partitioned []int){

	pivot := unsorted[0]

	var left, mid, right []int

	for _, item := range unsorted {

		if item < pivot {

			left = append(left, item)

		} else if item > pivot {

			right = append(right, item)

		} else {

			mid = append(mid, item)

		}
	}

	partitioned = append(partitioned, left...)

	partitioned = append(partitioned, mid...)

	partitioned = append(partitioned, right...)

	return

}