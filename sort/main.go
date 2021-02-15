package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 常数
const (
	Size     = 20
	NumRange = 100
)

func main() {

	testArray := make([]int, Size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Size; i++ {
		testArray[i] = rand.Intn(NumRange)
	}
	if len(testArray) <= 1 {
		return
	}

	fmt.Println(testArray)

	// Test QuickSort Success
	// QuickSort(testArray, 0, len(testArray)-1)

	// Test MergeSort Success
	MergeSort(testArray, 0, len(testArray)-1)

	// Test BubbleSort Success
	// BubbleSort(testArray)

	// Test InsertionSort Success
	// InsertionSortV2(testArray)

	// Test SelectionSort Success
	// SelectionSort(testArray)
	fmt.Println(testArray)

}
