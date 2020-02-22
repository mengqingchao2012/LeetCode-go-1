package main

import . "LeetCode-go/utils"

//解法一：自顶向下，遍历所有结点，判断是否满足平衡二叉树
//Time：O(nlogn)，Space：O(n)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(getHeight(root.Left)-getHeight(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)

	return max(left, right) + 1
}

//解法二：自底向上，遍历结点的同时，发现任意节点不是平衡二叉树，直接返回
//Time：O(n)，Space：O(n)
func isBalanced1(root *TreeNode) bool {
	return getHeightAndCheck(root) != -1
}

func getHeightAndCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeightAndCheck(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeightAndCheck(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}
