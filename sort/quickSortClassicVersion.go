package main

import (
	"math/rand"
)

/**
经典快速排序，固定基准数
随机快排
*/
const (
	TailPivot = 0
	HeadPivot = 1
	RandPivot = 2
)

// QuickSort is
func QuickSort(a []int, head, tail int) {
	if head >= tail {
		return
	}
	// pivot := partition(a, head, tail, TailPivot)
	// pivot := partition(a, head, tail, HeadPivot)
	pivot := partition(a, head, tail, RandPivot)
	QuickSort(a, head, pivot-1)
	QuickSort(a, pivot+1, tail)
}

func partition(a []int, head, tail, Flag int) int {
	switch Flag {
	case 0:
		return partitionWithTailPivot(a, head, tail)
	case 1:
		return partitionWithHeadPivot(a, head, tail)
	case 2:
		return partitionWithRandomPivot(a, head, tail)
	default:
		panic("no support for Flag")
	}
}

func partitionWithTailPivot(a []int, head, tail int) int {
	key := a[tail]
	i, j := head, tail
	for i != j {
		for i != j && a[i] <= key {
			i++
		}
		if i != j {
			a[j] = a[i]
			j--
		}
		for i != j && a[j] >= key {
			j--
		}
		if i != j {
			a[i] = a[j]
			i++
		}
	}
	a[i] = key
	return i
}

func partitionWithHeadPivot(a []int, head, tail int) int {
	key := a[head]
	i, j := head, tail
	for i != j {
		for i != j && a[j] >= key {
			j--
		}
		if i != j {
			a[i] = a[j]
			i++
		}
		for i != j && a[i] <= key {
			i++
		}
		if i != j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = key
	return i

}

func partitionWithRandomPivot(a []int, head, tail int) int {
	pivot := rand.Intn(tail-head+1) + head
	a[pivot], a[tail] = a[tail], a[pivot]
	return partitionWithTailPivot(a, head, tail)
}
