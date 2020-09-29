package main

import ."LeetCode-go/utils"

func sumNumbers(root *TreeNode) int {
	return search(root, 0)
}

func search(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = sum * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return search(root.Left, sum) + search(root.Right, sum)
}
