package main

// 方法一：dp
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	// dp[i] 表示以第 i 个数字结尾的最长增长子序列的长度
	dp := make([]int, n)
	// 将 dp 中的每个元素都更新为 1，因为单个字符的最长递增子序列就是它本身
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	res := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] <= dp[j] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

// 方法二：贪心 + 二分
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	subq := make([]int, n + 1) // 表示最终生成的最长递增子序列
	maxLen := 0 // 表示递增子序列的长度
	for _, v := range nums {
		l, r := 0, maxLen
		for l < r { // 注意：是在 subq 中进行二分
			mid := (l + r + 1) >> 1
			if subq[mid] < v {
				l = mid
			} else {
				r = mid - 1
			}
		}
		maxLen = Max(maxLen, r + 1)
		subq[r + 1] = v
	}
	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
