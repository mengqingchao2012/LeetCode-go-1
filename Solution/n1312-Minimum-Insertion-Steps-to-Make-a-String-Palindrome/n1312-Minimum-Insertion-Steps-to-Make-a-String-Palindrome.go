package main

import ."LeetCode-go/utils"

func minInsertions(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	maxLen := make([][]int, n)
	for i := 0; i < n; i++ {
		maxLen[i] = make([]int, n)
		maxLen[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				maxLen[i][j] = 2 + maxLen[i + 1][j - 1]
			} else {
				maxLen[i][j] = Max(maxLen[i + 1][j], maxLen[i][j - 1])
			}
		}
	}
	return n - maxLen[0][n - 1]
}
