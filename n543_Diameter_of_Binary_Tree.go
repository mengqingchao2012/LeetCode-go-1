package main

import . "LeetCode-go/utils"

//Time：O(n)，Space：O(n)
//自底向上返回节点的深度，所以只需遍历所有节点一次，时间复杂度为 O(n)
func diameterOfBinaryTree(root *TreeNode) int {
	depth := 0
	maxTreeDepth(root, &depth)
	return depth
}

func maxTreeDepth(root *TreeNode, depth *int) int {
	if root == nil {
		return 0
	}

	left := maxTreeDepth(root.Left, depth)
	right := maxTreeDepth(root.Right, depth)
	*depth = Max(*depth, left+right)
	return Max(left, right) + 1
}
