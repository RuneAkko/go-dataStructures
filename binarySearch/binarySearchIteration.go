package main

// 不重复，正序，查找等于
func bSearchWithIteration(a []int, value int) (bool, int) {
	start, end := 0, len(a)-1
	for start >= end {
		mid := start + (end-start)/2
		if a[mid] == value {
			return true, mid
		}
		if a[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false, -1
}

// 重复，正序数组，查找第一个
func bSearchFirstWithIteration(a []int, value int) (bool, int) {
	start, end := 0, len(a)-1
	for start <= end {
		mid := start + (end-start)/2
		if a[mid] == value {
			if mid == 0 || a[mid-1] != value {
				return true, mid
			}
			end = mid - 1
		}
		if a[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false, -1
}

// 重复，正序数组，查找最后一个
func bSearchLastWithIteration(a []int, value int) (bool, int) {
	start, end := 0, len(a)-1
	for start <= end {
		mid := start + (end-start)/2
		if a[mid] == value {
			if mid == len(a)-1 || a[mid+1] != value {
				return true, mid
			}
			start = mid + 1
		}
		if a[mid] > value {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false, -1
}

// 重复，正序，数组，查找第一个，大于等于
func bSearchFirstGreaterOrEqualWithIteration(a []int, value int) (bool, int) {
	start, end := 0, len(a)-1
	for start <= end {
		mid := start + (end-start)/2
		if a[mid] >= value {
			if mid == 0 || a[mid-1] < value {
				return true, mid
			}
			end = mid - 1
		}
		start = mid + 1
	}
	return false, -1
}
