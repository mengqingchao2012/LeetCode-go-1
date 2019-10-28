package main

func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	low, high := 0, size - 1

	for low <= high {
		mid := low + (high -low >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
