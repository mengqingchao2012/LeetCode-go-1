package main

func searchInsert(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}

	low, high := 0, size - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low // 注意插入位置返回的是 low
}