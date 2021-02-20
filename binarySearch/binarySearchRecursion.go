package main

// 不重复，正序，查找等于
func bSearchWithRecursion(a []int, start, end, value int) (bool, int) {
	if start > end {
		return false, -1
	}
	mid := start + (end-start)/2
	if a[mid] == value {
		return true, mid
	}
	if a[mid] > value {
		return bSearchWithRecursion(a, start, mid-1, value)
	}
	return bSearchWithRecursion(a, mid+1, end, value)
}

// 重复，正序数组，查找第一个出现的value
func bSearchFirstWithRecursion(a []int, start, end, value int) (bool, int) {
	if start > end {
		return false, -1
	}
	mid := start + (end-start)/2
	if a[mid] == value {
		if mid == 0 || a[mid-1] != value {
			return true, mid
		}
		return bSearchFirstWithRecursion(a, start, mid-1, value)
	}
	if a[mid] > value {
		return bSearchWithRecursion(a, start, mid-1, value)
	}
	return bSearchWithRecursion(a, mid+1, end, value)
}

// 重复，正序，数组，查找最后一个
func bSearchLastWithRecursion(a []int, start, end, value int) (bool, int) {
	if start > end {
		return false, -1
	}
	mid := start + (end-start)/2
	if a[mid] == value {
		if mid == len(a)-1 || a[mid+1] != value {
			return true, mid
		}
		return bSearchFirstWithRecursion(a, mid+1, end, value)
	}
	if a[mid] > value {
		return bSearchWithRecursion(a, start, mid-1, value)
	}
	return bSearchWithRecursion(a, mid+1, end, value)
}

// 重复，正序，数组，查找第一个，大于等于
func bSearchFirstGreaterOrEqualWithRecursion(a []int, start, end, value int) (bool, int) {
	if start > end {
		return false, -1
	}
	mid := start + (end-start)/2
	if a[mid] >= value {
		if mid == 0 || a[mid-1] < value {
			return true, mid
		}
		return bSearchFirstWithRecursion(a, start, mid-1, value)
	}
	return bSearchWithRecursion(a, mid+1, end, value)
}
