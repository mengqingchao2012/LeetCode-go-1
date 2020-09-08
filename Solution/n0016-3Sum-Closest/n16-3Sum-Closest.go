package main

import (
	. "LeetCode-go/utils"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	small := math.MaxInt64

	for i := len(nums) - 1; i >= 2; i-- {
		left, right := 0, i - 1
		for left < right {
			diff := target - nums[left] - nums[right] - nums[i]

			if diff == 0 {
				return target
			}

			if Abs(diff) < Abs(small) {
				small = diff
			}

			if diff > 0 {
				left++
			} else {
				right--
			}
		}
	}
	return target - small
}
