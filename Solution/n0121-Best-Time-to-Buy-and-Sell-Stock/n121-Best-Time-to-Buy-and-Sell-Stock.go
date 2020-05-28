package main

import . "LeetCode-go/utils"

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
