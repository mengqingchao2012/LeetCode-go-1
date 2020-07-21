package main

func findLength(A []int, B []int) int {
	ans := 0
	n, m := len(A), len(B)
	dp := make([][]int, n + 1)
	for i := 0; i < n + 1; i++ {
		dp[i] = make([]int, m + 1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i + 1][j + 1] + 1
			} else {
				dp[i][j] = 0
			}

			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}