package main

import "sort"

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	count := 0
	for i := len(nums) - 1; i >= 2; i-- {
		left, right := 0, i - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if target > sum {
				// 因为数组有序，则如果当前 sum < target，之后 left - right 之间的任意下标对应的值都满足条件
				count += right - left
				left++
			} else {
				right--
			}
		}
	}
	return count
}
