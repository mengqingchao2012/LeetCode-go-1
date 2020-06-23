package main

import "LeetCode-go/utils"

func reachNumber(target int) int {
	t := utils.Abs64(int64(target))
	var sum, n int64
	for sum < t || ((sum-t)&1) == 1 {
		n++
		sum += n
	}
	return int(n)
}
