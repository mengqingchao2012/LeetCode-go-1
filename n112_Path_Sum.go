package main

import ."LeetCode-go/utils"

//递归法：Time：O(n)，Space：O(n)
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil { //叶子节点，值要等于sum
		return root.Val == sum
	}

	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}

//迭代法：Time：O(n)，Space：O(n)
func hasPathSum1(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	//使用双栈，一个用来存储节点，另一个用来存储sum - val后的值
	var stack []*TreeNode
	var sumStack []int

	stack = append(stack, root)
	sumStack = append(sumStack, sum)

	for len(stack) != 0 {
		top := len(stack) - 1
		node := stack[top]
		s := sumStack[top]
		if node.Left == nil && node.Right == nil && node.Val == s {
			return true
		}
		stack = stack[: top]
		sumStack = sumStack[: top]
		temp := s - node.Val
		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, temp)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, temp)
		}
	}
	return false
}