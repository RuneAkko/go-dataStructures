package main

/**
归并排序
递归实现
正序
*/

// 最小最大
const (
	INTMAX = int(^uint(0) >> 1)
)

// MergeSort is
func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	MergeSort(arr, l, m)
	MergeSort(arr, m+1, r)
	merge(arr, l, m, m+1, r)
	mergeWithSentail(arr, l, m, m+1, r)
}

func merge(arr []int, l, m, m1, r int) {
	res := make([]int, 0)
	i, j := l, m1
	for i <= m && j <= r {
		if arr[i] < arr[j] {
			res = append(res, arr[i])
			i++
		} else {
			res = append(res, arr[j])
			j++
		}
	}
	if i > m {
		res = append(res, arr[j:r+1]...)
	}
	if j > r {
		res = append(res, arr[i:m+1]...)
	}
	for index, ele := range res {
		arr[l+index] = ele
	}
}

func mergeWithSentail(arr []int, l, m, m1, r int) {
	tmpLeft := make([]int, m-l+1)
	tmpRight := make([]int, r-m1+1)
	copy(tmpLeft, arr[l:m+1])
	copy(tmpRight, arr[m1:r+1])
	tmpLeft = append(tmpLeft, INTMAX)
	tmpRight = append(tmpRight, INTMAX)
	res := make([]int, 0)
	for i, j, k := 0, 0, l; k <= r; k++ {
		if tmpLeft[i] < tmpRight[j] {
			res = append(res, tmpLeft[i])
			i++
		} else {
			res = append(res, tmpRight[j])
			j++
		}
	}
	for index, ele := range res {
		arr[l+index] = ele
	}
}
