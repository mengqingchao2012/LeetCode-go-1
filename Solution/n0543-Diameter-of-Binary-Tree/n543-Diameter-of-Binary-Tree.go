package main

import . "LeetCode-go/utils"

func diameterOfBinaryTree(root *TreeNode) int {
	depth := 0
	maxDepth(root, &depth)
	return depth
}

func maxDepth(root *TreeNode, depth *int) int {
	if root == nil { //递归终止条件
		return 0
	}
	leftDepth := maxDepth(root.Left, depth)
	rightDepth := maxDepth(root.Right, depth)
	*depth = Max(*depth, leftDepth+rightDepth)
	return Max(leftDepth, rightDepth) + 1
}
