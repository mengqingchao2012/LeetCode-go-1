package main

import ."LeetCode-go/utils"

// 方法一：
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Abs(getHeight(root.Left) - getHeight(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return Max(getHeight(root.Left), getHeight(root.Right)) + 1
}

// 方法二：
func isBalanced1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getHeightAndCheck(root) != -1
}

// 传入一个结点，检查其是否是平衡二叉树，若是则返回其高度，
// 若不是则返回 -1
func getHeightAndCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeightAndCheck(root.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := getHeightAndCheck(root.Right)
	if rightHeight == -1 {
		return -1
	}

	if Abs(leftHeight - rightHeight) > 1 {
		return -1
	}

	return Max(leftHeight, rightHeight) + 1
}
