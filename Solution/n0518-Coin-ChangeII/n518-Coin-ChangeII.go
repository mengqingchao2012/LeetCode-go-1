package main

import "sort"

// 方法一：递归，会超时
func change(amount int, coins []int) int {
	sort.Ints(coins)
	return count_partitions(amount, coins)
}

func count_partitions(amount int, coins []int) int {
	if amount == 0 {
		return 1
	}

	if len(coins) == 0 || amount < coins[0] {
		return 0
	}

	with_coin0 := count_partitions(amount - coins[0], coins)
	without_coin0 := count_partitions(amount, coins[1:])
	return with_coin0 + without_coin0
}

// 方法二：动态规划
func change1(amount int, coins []int) int {
	dp := make([]int, amount + 1) //dp[i] 指的是要凑成金额为i的情况，可以实现的组合的个数

	dp[0] = 1 // 表示要凑成金额为0的情况，只有一种实现方式，就是所有的硬币类型都取0个
	for _, coin := range coins {
		//金额 i 小于货币面额时，i - coin 为负，这种情况下不满足条件，因此直接令金额从面额值开始
		for i := coin; i <= amount; i++ {
			//状态转移方程：凑成金额为 i 的情况的组合数=当前轮不使用面额为coin的货币的所有组合数 + 使用了面额为coin的货币的组合数
				dp[i] = dp[i] + dp[i - coin]
		}
	}
	return dp[amount]
}
