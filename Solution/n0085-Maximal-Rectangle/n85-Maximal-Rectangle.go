package main

import ."LeetCode-go/utils"

func maximalRectangle(matrix [][]byte) int {
	row := len(matrix)
	if row == 0 {
		return 0
	}
	col := len(matrix[0])

	// 以任一行作为矩形的下边界，求出该行中各列的高度，这样就可以将问题变为第 84 题来求解
	heights := make([][]int, row)
	for i := 0; i < row; i++ {
		heights[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				if i > 0 {
					heights[i][j] = heights[i - 1][j] + 1
				} else {
					heights[i][j] = 1
				}
			}
		}
	}

	res := 0
	for i := 0; i < row; i++ {
		res = Max(res, largestRectangleArea(heights[i]))
	}
	return res
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	stack := make([]int, 0, n)
	left := make([]int, n)
	right := make([]int, n)

	walk := func (i int, bound int, res []int) {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1]
		} else {
			res[i] = bound
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		walk(i, -1, left)
	}

	stack = make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		walk(i, n, right)
	}

	max := 0
	for i := 0; i < n; i++ {
		max = Max(max, heights[i]*(right[i]-left[i]-1))
	}
	return max
}
