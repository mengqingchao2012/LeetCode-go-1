package main

import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s) + 1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		num, _ := strconv.Atoi(s[i - 1 : i])
		if num >= 1 && num <= 9 {
			dp[i] += dp[i - 1]
		}
		num, _ = strconv.Atoi(s[i - 2 : i])
		if num >= 10 && num <= 26 {
			dp[i] += dp[i - 2]
		}
	}
	return dp[len(s)]

}