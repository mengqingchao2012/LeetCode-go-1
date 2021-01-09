package main

import (
	. "LeetCode-go/utils"
)

func maxProfit(prices []int) int {
	if len(prices) < 2 { // 股票需要一天买入一天卖出，所以如果买卖天数小于2，则无法完成交易
		return 0
	}

	buy, profile := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
		} else {
			profile = Max(profile, prices[i]-buy)
		}
	}
	return profile
}

// dp
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][0] 表示第 i 天手上不持有股票时的最大利润
	// dp[i][1] 表示第 i 天手上持有股票时的最大利润
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [2]int{}
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i - 1][0], dp[i - 1][1] + prices[i])
		// 注意这里不要写成 Max(dp[i - 1][1], dp[i - 1][0] - prices[i])
		// 因为只能买卖一次，则卖出当天的利润就只能是买入价格的负值
		dp[i][1] = Max(dp[i - 1][1], - prices[i])
	}

	return dp[n - 1][0]
}
