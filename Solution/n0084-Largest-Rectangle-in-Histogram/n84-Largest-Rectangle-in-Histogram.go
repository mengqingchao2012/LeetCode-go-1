package main

import ."LeetCode-go/utils"

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
		max = Max(max, heights[i] * (right[i] - left[i] - 1))
	}
	return max
}
