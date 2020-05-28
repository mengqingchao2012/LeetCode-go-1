package main

import . "LeetCode-go/utils"

// 方法一：Time：O(nlogn），Space：O(n)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	} // 退递归条件，root==nil时，返回true
	leftValid := root.Left == nil || root.Val > maxVal(root.Left)                     // 验证root的值大于左子树的最大节点值
	rightValid := root.Right == nil || root.Val < minVal(root.Right)                  // 验证root的值小于右子树的最小节点值
	return leftValid && rightValid && isValidBST(root.Left) && isValidBST(root.Right) // 递归检查左子树和右子树
}

// 求出以当前节点为根节点，其子树中的最大值（最右叶节点的值）
func maxVal(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}

// 求出以当前节点为根节点，其子树中的最小值（最左叶节点的值）
func minVal(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

// 方法二：Time：O(n），Space：O(n)
func isValidBST1(root *TreeNode) bool {
	return checkBound(root, nil, nil)
}

func checkBound(root, lower, upper *TreeNode) bool {
	if root == nil {
		return true
	}
	if lower != nil && lower.Val >= root.Val {
		return false
	} // 注意不要漏了等号
	if upper != nil && upper.Val <= root.Val {
		return false
	} // 注意不要漏了等号
	return checkBound(root.Left, lower, root) && checkBound(root.Right, root, upper)
}
