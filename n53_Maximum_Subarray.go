package main

import (
	. "LeetCode-go/utils"
	"math"
)

//Time：O(n), Space:(1)
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	max, cur := math.MinInt64, 0
	for _, num := range nums {
		if cur <= 0 {
			cur = num
		} else {
			cur += num
		}
		max = Max(max, cur)
	}
	return max
}
