package main

func longestRepeatingSubstring(S string) int {
	n := len(S)
	if n == 0 {
		return 0
	}

	// dp[i][j] 表示 [i:] 范围内最长重复子串的长度，dp[i][i] = 0
	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n + 1)
	}

	maxLen := 0
	// 对于长度为 k 的子串，如果他是重复子串，则长度为 k - 1 的子串也是重复子串
	// i 是区间的起点，j 从 i + 1 开始遍历，如果两个指针指向的元素相同，则更新其长度为 dp[i - 1][j - 1] + 1
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if S[i - 1] == S[j - 1] {
				dp[i][j] = dp[i - 1][j - 1] + 1
			}
			if maxLen < dp[i][j] {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}
