package main

import ."LeetCode-go/utils"

// 方法一：前后缀分离
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// 为了方便计算，设置下标从 1 开始，则后续所有对 prices 的操作下标都要减 1
	mp := make([]int, n + 1)
	buy := prices[0] // 记录买入时的价格
	for i := 1; i <= n; i++ { // 遍历卖出时的价格
		mp[i] = Max(mp[i - 1], prices[i - 1] - buy)
		buy = Min(buy, prices[i - 1])
	}

	res := 0
	sell := 0 // 记录卖出时的价格
	for i := n; i > 0; i-- { // 反向遍历买入时的价格
		res = Max(res, sell - prices[i - 1] + mp[i - 1])
		sell = Max(sell, prices[i - 1])
	}
	return res
}

// 方法二：动态规划
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][0][k]：第 i 天，手上没有股票，已经进行了 k 次交易的最大利润
	// dp[i][1][k]：第 i 天，手上有股票，已经进行了 k 次交易的最大利润
	dp := make([][2][3]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2][3]int{}
	}

	// 注意：题目是最多进行两次交易，即可以进行交易的次数为 0 次，1 次，2 次
	dp[0][0][0] = 0
	dp[0][1][0] = -prices[0]
	dp[0][0][1] = 0
	dp[0][1][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0][0] = dp[i - 1][0][0]
		dp[i][1][0] = Max(dp[i - 1][1][0], dp[i - 1][0][0] - prices[i])
		dp[i][0][1] = Max(dp[i - 1][0][1], dp[i - 1][1][0] + prices[i])
		dp[i][1][1] = Max(dp[i - 1][1][1], dp[i - 1][0][1] - prices[i])
		dp[i][0][2] = Max(dp[i - 1][0][2], dp[i - 1][1][1] + prices[i])
	}
	return dp[n - 1][0][2]
}