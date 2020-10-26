package main

import ."LeetCode-go/utils"

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	dp := make([]int, n) // 表示抢劫前 i 所房子得到的最大收益
	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		// 对于第 i 所房子，如果选择不抢，则收益等于 dp[i - 1]，若选择抢，则收益
		// 等于 dp[i - 2] + nums[i]
		dp[i] = Max(dp[i - 1], dp[i - 2] + nums[i])
	}
	return dp[n - 1]
}

// 空间优化
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// prev1 表示上一次的最大利润，prev2 表示上上次的最大利润
	prev1, prev2 := 0, 0

	for i := 0; i < n; i++ {
		prev2, prev1 = prev1, Max(prev1, prev2 + nums[i])
	}
	return prev1
}
