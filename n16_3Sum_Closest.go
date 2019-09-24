package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, min := 0, math.MaxInt64

	for i := len(nums) - 1; i >= 2; i-- {
		j, k := 0, i - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			} else if sum < target {
				j++
			} else {
				k--
			}

			diff := differ(sum, target)

			if diff < min {
				min = diff
				res = sum
			}
		}
	}
	return res
}


