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

// dp泛化版本
func maxProfit1(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// MP 表示最大利润，是一个二维数组
	// 第一维表示天数
	// 第二维表示当天执行的操作：什么都不做是 0，买入是 1，卖出是 2
	MP := make([][]int, n)
	MP[0] = make([]int, 3)
	// 初始化边界条件：对于第一天，什么都不做和卖出利润都设为 0，买入的利润是负的
	MP[0][0], MP[0][1], MP[0][2] = 0, -prices[0], 0
	res := 0

	for i := 1; i < n; i++ {
		MP[i] = make([]int, 3)
		// 第 i 天状态是没有交易，则利润等于第 i - 1 天没有交易的利润
		MP[i][0] = MP[i-1][0]
		// 第 i 天状态是买入，有两种情况：
		// 第 i - 1 天买入的，第 i 天什么都没做
		// 第 i - 1 天没有买入，第 i 天买入的
		// 利润等于两者中的最大值
		MP[i][1] = MultiMax(MP[i - 1][1], MP[i - 1][0] - prices[i])
		// 第 i 天状态变为卖出，则只能是第 i - 1 天已经是买入状态，并在第 i 天卖出
		MP[i][2] = MP[i - 1][1] + prices[i]
		res = MultiMax(res, MP[i][0], MP[i][1], MP[i][2])
	}
	return res
}
