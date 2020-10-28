package main

func countSubstrings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	cnt := n // 注意这里 cnt 要初始化为 n，因为单个字符都是回文子串
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i + 1 == j {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = s[i] == s[j] && dp[i + 1][j - 1]
			}

			if dp[i][j] {
				cnt++
			}
		}
	}
	return cnt
}

func countSubstrings1(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	res := 0
	for i := 0; i < n; i++ {
		res += expend(s, i, i)
		res += expend(s, i, i + 1)
	}
	return res
}

func expend(s string, left, right int) int {
	cnt, size := 0, len(s)
	for left >= 0 && right < size && s[left] == s[right] {
		cnt++
		left--
		right++
	}
	return cnt
}
