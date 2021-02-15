package main

/**
选择排序
正序
*/

// SelectionSort is
func SelectionSort(a []int) {
	var min, minI int
	for i := 0; i < len(a); i++ {
		min = a[i]
		minI = i
		for j := i; j < len(a); j++ {
			if a[j] < min {
				min = a[j]
				minI = j
			}
		}
		a[i], a[minI] = a[minI], a[i]
	}
}
