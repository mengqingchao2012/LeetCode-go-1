package main

import ."LeetCode-go/utils"

func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	if n1 == 0 || n2 == 0 {
		return 0
	}

	// dp[i][j] 表示在 text1 中取 i 个字符，text2 中取 j 个字符能够组成的最大公共子串
	// 注意：如果是从 text1 或 text2 中取出的字符为 0，则无法构成公共子串，故 dp[i][0] 和 dp[0][j] 都为 0
	dp := make([][]int, n1 + 1)
	for i := 0; i < n1 + 1; i++ {
		dp[i] = make([]int, n2 + 1)
	}

	maxLen := 0
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			// 如果当前下标对应的两个字符相同，则最长子串等于两个下标同时回退一格时的长度加 1
			if text1[i - 1] == text2[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			} else {
				// 如果不相同，则说明最长子串中 i 和 j 对应的元素要么只有 1 个在结果中，要么两个都不在，
				// 因此最大长度等于跳过 i（j在结果中）或者跳过 j（i在结果中）能得到的最大长度
				dp[i][j] = Max(dp[i - 1][j], dp[i][j - 1])
			}
			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}