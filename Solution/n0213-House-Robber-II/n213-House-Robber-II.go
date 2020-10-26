package main

import ."LeetCode-go/utils"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return Max(robMax(nums[1:]), robMax(nums[:n - 1]))
}

func robMax(nums []int) int {
	n := len(nums)
	prev2, prev1 := 0, 0

	for i := 0; i < n; i++ {
		prev2, prev1 = prev1, Max(prev1, prev2 + nums[i])
	}
	return prev1
}
