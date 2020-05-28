package main

import "math"

func findPeakElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	low, high := 0, length
	left, right := 0, 0
	for low < high { // 注意条件是 low < high
		mid := low + ((high - low) >> 1)
		if mid-1 < 0 {
			left = math.MinInt64
		} else {
			left = nums[mid-1]
		}
		if mid+1 >= length {
			right = math.MinInt64
		} else {
			right = nums[mid+1]
		}

		if nums[mid] > left && nums[mid] > right {
			return mid
		} else if nums[mid] < left {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low // 退出循环时，low=high，随便返回哪个值都可以
}
