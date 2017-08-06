package merge_sort

func Merge(a []int, b []int) (c []int){

	for {

		if len(a) == 0 {
			c = append(c, b...)
			return c

		} else if len(b) == 0 {

			c = append(c, a...)
			return c

		} else if a[0] <= b[0] {

			c = append(c, a[0])
			a = a[1:]

		} else {

			c = append(c, b[0])
			b = b[1:]

			}

	}


}

func Sort(a []int) (b []int) {

	if len(a) == 1 || isOrdered(a) {

		return a

	} else if len(a) == 2 {

		{

			return []int{a[1], a[0]}
		}

	} else {

		return Merge(Sort(a[len(a)/2:]), Sort(a[:len(a)/2]))

	}

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
