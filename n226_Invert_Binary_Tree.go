package main

import ."LeetCode-go/utils"

//递归法：要遍历所有节点一次，时间复杂度为O(n)，递归深度为n
//Time：O(n)，Space：O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

//迭代法：Time：O(n)，Space：O(n)
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var stack []*TreeNode //使用辅助栈或辅助队列
	stack = append(stack, root)

	for len(stack) != 0 {
		//出栈
		length := len(stack) - 1
		node := stack[length]
		stack = stack[:length]
		//翻转
		node.Left, node.Right = node.Right, node.Left
		//子节点再入栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return root
}