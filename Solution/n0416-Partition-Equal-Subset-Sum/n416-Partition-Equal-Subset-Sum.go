package main

import "math"

// 方法一：二维数组
func canPartition(nums []int) bool {
	n := len(nums)
	// 数组元素小于 2 个，直接返回 false
	if n < 2 {
		return false
	}
	sum, max := 0, math.MinInt32
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	if sum % 2 != 0 {
		return false
	}
	sum /= 2
	// 如果数组中的最大元素比数组元素和的一半还大，则剩余的元素和肯定
	// 比元素和的一半小，返回 false
	if sum < max {
		return false
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum + 1)
		dp[i][0] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j < sum + 1; j++ {
			if nums[i] <= j {
				dp[i][j] = dp[i - 1][j] || dp[i - 1][j - nums[i]]
			} else {
				dp[i][j] = dp[i - 1][j]
			}
		}
	}
	return dp[n - 1][sum]
}

// 方法二：一维数组
func canPartition1(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, math.MinInt32
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	if sum % 2 != 0 {
		return false
	}
	sum /= 2
	if sum < max {
		return false
	}

	dp := make([]bool, sum + 1)
	dp[0] = true
	for i := 1; i < n; i++ {
		// 注意此时 j 要从后向前遍历，同时可以将 j >= nums[i] 的条件提到
		// for 的状态语句中
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j - nums[i]]
		}
	}
	return dp[sum]
}
