package main

func tribonacci(n int) int {
	dp := make([]int, n + 1)
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	}

	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2] + dp[i - 3]
	}
	return dp[n]
}

func tribonacci1(n int) int {
	if n == 0 {
		return 0
	} else if n < 3 {
		return 1
	}

	a, b, c := 0, 1, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a + b + c
	}
	return c
}
