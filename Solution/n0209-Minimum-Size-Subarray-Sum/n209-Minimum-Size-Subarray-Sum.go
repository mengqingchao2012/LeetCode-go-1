package main

import ."LeetCode-go/utils"

import "math"

func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	winSum, start, size := 0, 0, math.MaxInt64
	for end := 0; end < length; end++ {
		winSum += nums[end]
		for winSum >= s {
			size = Min(size, end - start + 1)
			winSum -= nums[start]
			start += 1
		}
	}
	if size == math.MaxInt64 {
		size = 0
	}
	return size
}
