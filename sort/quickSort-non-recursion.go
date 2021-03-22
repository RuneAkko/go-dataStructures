package main

func QuickSortIteration(a []int) {
	if len(a) < 2 {
		return
	}

	s := make([]int, 0)
	s = append(s, 0)
	s = append(s, len(a)-1)

	for len(s) != 0 {

		tail := s[len(s)-1]
		s = s[:len(s)-1]

		head := s[len(s)-1]
		s = s[:len(s)-1]

		pivot := partitionWithHeadPivot(a, head, tail)

		if pivot-1 > head {
			s = append(s, head)
			s = append(s, pivot-1)
		}

		if pivot+1 < tail {
			s = append(s, pivot+1)
			s = append(s, tail)
		}

	}

}
