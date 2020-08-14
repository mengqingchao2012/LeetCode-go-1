package main

import ."LeetCode-go/utils"

func maxProduct(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}

	gmax, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < n; i++ {
		a, b, c := curMax * nums[i], curMin * nums[i], nums[i]
		curMax = MultiMax(a, b, c)
		curMin = MultiMin(a, b, c)
		gmax = MultiMax(gmax, curMax)
	}

	return gmax
}
