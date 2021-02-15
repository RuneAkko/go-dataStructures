package main

/**
冒泡排序 正序
*/

// BubbleSort is
func BubbleSort(a []int) {
	flag := false
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}
