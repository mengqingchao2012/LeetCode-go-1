package main

import ."LeetCode-go/utils"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	la, lb := len(A), len(B)
	if la == 0 || lb == 0 {
		return [][]int{}
	}

	res := [][]int{}
	i, j := 0, 0

	for i < la && j < lb {
		start := Max(A[i][0], B[j][0])
		end := Min(A[i][1], B[j][1])
		if start <= end {
			res = append(res, []int{start, end})
		}

		if A[i][1] < B[j][1] {
			i += 1
		} else {
			j += 1
		}
	}
	return res
}