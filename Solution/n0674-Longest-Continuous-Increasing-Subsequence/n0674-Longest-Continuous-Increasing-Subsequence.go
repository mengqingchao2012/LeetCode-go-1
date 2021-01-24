package main

import ."LeetCode-go/utils"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	res := 0
	slow, fast := 0, 1
	for fast < n {
		if nums[fast] <= nums[fast - 1] {
			res = Max(res, fast - slow)
			slow = fast
		}
		fast++
	}
	res = Max(res, fast - slow)
	return res
}