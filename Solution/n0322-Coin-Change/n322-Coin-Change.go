package main

import . "LeetCode-go/utils"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp[i]表示要凑成金额为 i 的情况，需要的最小硬币数
	dp[0] = 0                   //总金额为0时，不管使用什么硬币都无法构成有效解，所以dp[0]=0，表示需要的硬币数为0

	for i := 1; i <= amount; i++ {
		// 金额从1开始涨到amount，对应的硬币数应该设为一个很大的值，因为不确定到底需要多少个硬币才能满足条件，甚至有可能是
		// 现有的硬币都无法满足条件，此时按题意应该设置为 -1，但是由于我们要求最少的硬币数，如果设为 -1，则后续的取最小值的
		// 结果都是 -1，不满足条件，故改为设为一个非常大的值，此处直接使用 amount + 1来表示，免于使用 math.MaxInt，同时也避免了
		// 后续的 +1 操作导致溢出的情况
		dp[i] = amount + 1
	}

	for _, coin := range coins {
		//当金额 i 小于硬币面额 coin时，i - coin 是个负数，不满足条件，因此金额直接从选定硬币的面额开始
		for i := coin; i <= amount; i++ {
			// 如果dp[i - coin] 是取的最大值，则与 dp[i] 取小后结果还是 dp[i]（dp[i]也可能是最大值），可以直接略过这种情况，
			// 少一次取最小值的操作
			if dp[i-coin] != amount+1 {
				dp[i] = Min(dp[i], dp[i-coin]+1)
			}
		}
	}
	// 最后的判断，如果 dp[amount] = amount + 1，说明没有找到可用的解，返回 -1
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 使用二维数组
func coinChange1(coins []int, amount int) int {
	count := len(coins)
	// dp 表示使用前 i 种硬币凑总数 j 需要的最小硬币数
	// 注意这里第一维的范围是 [0,len(coins)]，0 表示使用前 0 种硬币，
	// 之所以这样写是为了简化边界条件的处理情况
	dp := make([][]int, count + 1)
	for i := range dp {
		dp[i] = make([]int, amount + 1)
	}
	for i := 1; i <= amount; i++ {
		dp[0][i] = 1e8
	}

	for i := 1; i <= count; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i - 1][j]
			if coins[i - 1] <= j {
				dp[i][j] = Min(dp[i][j], dp[i][j - coins[i - 1]] + 1)
			}
		}
	}

	if dp[count][amount] == 1e8 {
		return -1
	}
	return dp[count][amount]
}
