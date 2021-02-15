package main

/**
插入排序 正序
*/

// InsertionSort is
func InsertionSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				temp := a[i]
				for k := i; k > j; k-- {
					a[k] = a[k-1]
				}
				a[j] = temp
			}
		}
	}
}

// InsertionSortV2 is
func InsertionSortV2(a []int) {
	for i := 1; i < len(a); i++ {
		val := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > val {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = val
	}
}
