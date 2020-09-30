package main

import . "LeetCode-go/utils"

func diameterOfBinaryTree(root *TreeNode) int {
	depth := 0
	maxDepth(root, &depth)
	return depth
}

// maxDepth 求得的是节点的深度, depth 才是直径
func maxDepth(root *TreeNode, depth *int) int {
	if root == nil { //递归终止条件
		return 0
	}

	// 求出左右子树的深度
	leftDepth := maxDepth(root.Left, depth)
	rightDepth := maxDepth(root.Right, depth)
	// 将当前节点左右子树的深度相加，并将和与之前保存的最大直径相比较，取较大值
	*depth = Max(*depth, leftDepth+rightDepth)

	// 返回当前节点的深度，即左右子树中深度较大的子树深度加 1
	return Max(leftDepth, rightDepth) + 1
}
