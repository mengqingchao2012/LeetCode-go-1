package main

import ."LeetCode-go/utils"

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return cost[0]
	}

	first, second, cur := cost[0], cost[1], 0
	for i := 2; i < length; i++ {
		cur = Min(first, second) + cost[i]
		first = second
		second = cur
	}

	return Min(first, second)
}
