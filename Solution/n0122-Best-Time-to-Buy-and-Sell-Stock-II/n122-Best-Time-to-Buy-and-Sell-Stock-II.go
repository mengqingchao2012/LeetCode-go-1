package main

import ."LeetCode-go/utils"

// 动态规划
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// 第一维表示天数，第二维表示是否持有股票
	MP := make([][]int, n)
	MP[0] = make([]int, 2)

	MP[0][0], MP[0][1] = 0, -prices[0]
	res := 0

	for i := 1; i < n; i++ {
		MP[i] = make([]int, 2)

		MP[i][0] = MultiMax(MP[i - 1][0], MP[i - 1][1] + prices[i])
		MP[i][1] = MultiMax(MP[i - 1][0] - prices[i], MP[i - 1][1])

		res = MultiMax(res, MP[i][0], MP[i][1])
	}
	return res
}

// 贪心：只要后一天的价格高于当天价格，就执行一次买入卖出
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	res := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i - 1] {
			res += prices[i] - prices[i - 1]
		}
	}
	return res
}