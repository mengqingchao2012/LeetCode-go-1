package main

import ."LeetCode-go/utils"

func longestOnes(A []int, K int) int {
	start, maxCount, res := 0, 0, 0

	for end := 0; end < len(A); end++ {
		if A[end] == 1 {
			maxCount++
		}

		if end - start + 1 - maxCount > K {
			if A[start] == 1 {
				maxCount--
			}
			start++
		}
		res = Max(res, end - start + 1)
	}
	return res
}
