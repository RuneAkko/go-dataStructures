package main

/**
归并排序
递归实现
正序
*/

// 最小最大
const (
	INTMIN = int(^uint(0) >> 1)
	INTMAX = ^INTMIN
)

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, m+1, r)
}

func merge(arr []int, l, m, m1, r int) {
	res := make([]int, m+r-l-m1+2)
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
		arr = res
	}
	if j > r {
		res = append(res, arr[i:m1+1]...)
		arr = res
	}
}

func mergeWithSentail(arr []int, l, m, m1, r int) {

	tmpLeft := arr[l : m+1]
	tmpRight := arr[m1 : r+1]
	tmpLeft = append(tmpLeft, INTMAX)
	tmpRight = append(tmpRight, INTMAX)
	res := make([]int, len(tmpLeft)+len(tmpRight)-2)
	for k, i, j := l, 0, 0; k <= r; k++ {
		if tmpLeft[i] < tmpRight[j] {
			res[k] = tmpLeft[i]
			i++
		} else {
			res[k] = tmpRight[j]
			j++
		}
	}
	arr = res
}
