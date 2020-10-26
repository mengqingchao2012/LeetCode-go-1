package main

import ."LeetCode-go/utils"

func rob(root *TreeNode) int {
	res := robMax(root)
	return Max(res[0], res[1])
}

func robMax(root *TreeNode) [2]int {
	if root == nil {
		return [2]int{0, 0}
	}

	res := [2]int{0, 0}
	left := robMax(root.Left)
	right := robMax(root.Right)

	res[0] = Max(left[0], left[1]) + Max(right[0], right[1])
	res[1] = root.Val + left[0] + right[0]
	return res
}
