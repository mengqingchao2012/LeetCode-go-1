package main

func searchRange(nums []int, target int) []int {
	size := len(nums)
	if size == 0 {
		return []int{-1, -1}
	}

	end := findLastIndex(&nums, target)
	begin := findLastIndex(&nums, target - 1) + 1

	if begin >= 0 && begin <=end && end < size {
		return []int{begin, end}
	}

	return []int{-1, -1}
}

func findLastIndex(nums *[]int, target int) int {
	size := len(*nums)
	low, high := 0, size - 1

	for low <= high {
		mid := low + (high - low >> 1)
		if (*nums)[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}