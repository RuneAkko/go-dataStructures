package heap

import (
	"errors"
	"fmt"
)

// Heap 使用数组实现的堆（最大堆）
type Heap struct {
	arr  []int
	size int // 存储容量
	n    int // 当前存储元素个数,恰好是最末尾元素的索引
}

// Newheap return a new heap with capacity cap
func Newheap(cap int) *Heap {
	return &Heap{
		arr:  make([]int, cap+1),
		size: cap,
		n:    0,
	}
}

// BuildheapWithDown2Up 由输入数组建立一个最大堆,按自下而上建堆
func BuildheapWithDown2Up(arr []int) *Heap {
	if len(arr) < 1 {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("panic:%s\n", p)
			}
		}()
		panic(errors.New("input array is nil"))
	}
	heap := &Heap{
		arr:  make([]int, len(arr)+1),
		size: len(arr),
		n:    0,
	}
	for _, ele := range arr {
		heap.insert(ele)
	}
	return heap
}

// BuildheapWithUp2Down 由输入数组建立一个最大堆,按自上而下建堆
func BuildheapWithUp2Down(arr []int) *Heap {
	if len(arr) < 1 {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("panic:%s\n", p)
			}
		}()
		panic(errors.New("input array is nil"))
	}
	heap := &Heap{
		arr:  make([]int, len(arr)+1),
		size: len(arr),
		n:    0,
	}
	for i := 1; i <= heap.size; i++ {
		heap.arr[i] = arr[i-1]
	}
	heap.n = heap.size

	for i := heap.n / 2; i >= 1; i-- {
		heapify(heap, i)
	}
	return heap
}

// 堆化，从选中的start节点自上而下堆化
func heapify(h *Heap, start int) {
	for i := start; start*2 <= h.n; {
		if 2*i+1 <= h.n && h.arr[2*i+1] > h.arr[2*i] {
			h.arr[i], h.arr[2*i+1] = h.arr[2*i+1], h.arr[i]
		} else {
			h.arr[i], h.arr[2*i] = h.arr[2*i], h.arr[i]
		}
	}
}

// 堆排序
func heapSort(arr []int) []int {
	heap := BuildheapWithDown2Up(arr)
	// heap := BuildheapWithUp2Down(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], _ = heap.popTop()
	}
	return arr
}

// insert 下自上 堆化 插入新元素
func (h *Heap) insert(ele int) error {
	if h.n >= h.size {
		return errors.New("heap is full")
	}
	h.n++
	h.arr[h.n] = ele
	for i := h.n; i/2 > 0 && h.arr[i/2] < h.arr[i]; {
		h.arr[i], h.arr[i/2] = h.arr[i/2], h.arr[i]
		i = i / 2
	}
	return nil
}

// popTop 弹出堆顶元素，堆化剩余堆
func (h *Heap) popTop() (int, error) {
	if h.n < 1 {
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("panic:%s\n", p)
			}
		}()
		panic(errors.New("heap is empty"))
	}
	ele := h.arr[1]
	h.arr[1] = h.arr[h.n]
	h.n--
	for i := 1; 2*i <= h.n; {
		max := h.arr[2*i]
		if 2*i+1 <= h.n && h.arr[2*i+1] > max {
			h.arr[i], h.arr[2*i+1] = h.arr[2*i+1], h.arr[i]
			i = 2*i + 1
		} else {
			h.arr[i], h.arr[2*i] = h.arr[2*i], h.arr[i]
			i = 2 * i
		}
	}
	return ele, nil
}
