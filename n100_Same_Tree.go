package main

import (
	. "LeetCode-go/utils"
)

//递归解法：Time：O(n)，Space：O(n)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var res bool
	switch{
	case p == nil && q == nil: //如果两个根节点都为空，返回true
		res = true
	case p == nil || q == nil: //如果两个根节点一个为空另一个非空，返回false
		res = false
	default: //两个根节点都非空，比较根节点的值，并递归的判断其左右子树
		res = p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return res
}

//迭代法：Time：O(n)，Space：O(n)
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	var stack []*TreeNode
	stack = append(stack, p, q) //两个根节点入栈

	for len(stack) != 0 { //栈不为空时，将栈顶两个节点出栈
		node1 := stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		node2 := stack[len(stack) - 1]
		stack = stack[: len(stack) - 1]
		switch {
		case node1 == nil && node2 == nil: //如果两个节点均为空，则继续循环
			continue
		case node1 == nil || node2 == nil:
			return false
		case node1.Val != node2.Val:
			return false
		}

		stack = append(stack, node1.Left, node2.Left, node1.Right, node2.Right) //将两个节点的左右节点入栈
	}
	return true
}

