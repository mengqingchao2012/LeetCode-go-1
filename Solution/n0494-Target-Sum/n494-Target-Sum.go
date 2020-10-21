package main

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum < S || (sum + S) % 2 == 1 {
		return 0
	}

	n := len(nums)
	target := (sum + S) / 2
	// dp[i] 表示和为 i 时存在的所有解法
	dp := make([]int, target + 1)
	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j - nums[i]]
		}
	}
	return dp[target]
}
