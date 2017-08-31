package merge_sort

type MergeCounter struct {

	counter int
}

func (mc *MergeCounter) Merge(a []int, b []int) (c []int){

	if len(a) == 1 && len(b) == 1 && a[0] > b[0] {
		mc.counter = mc.counter +1
	}
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

			mc.counter = mc.counter +1
			c = append(c, b[0])
			b = b[1:]

			}

	}

}

func (mc *MergeCounter) Sort(a []int) (b []int) {

	if len(a) == 1 {

		return a

	} else {

		//fmt.Println(a[:len(a)/2 + len(a) % 2], a[len(a)/2 + len(a) % 2:])
		return mc.Merge(mc.Sort(a[:len(a)/2 + len(a) % 2]), mc.Sort(a[len(a)/2 + len(a) % 2:]))
	}

}

func CountInversions (a []int) (r int) {

	mc := MergeCounter{counter:0}

	mc.Sort(a)

	return mc.counter

}

