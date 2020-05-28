package main

import . "LeetCode-go/utils"

// 方法一：递归法
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}

// 方法二：迭代法
func rangeSumBST1(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	sum := 0
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil { // 注意这里先对取出来的节点值进行判断，将空节点先丢弃
			continue
		}
		if node.Val < L {
			stack = append(stack, node.Right)
		} else if node.Val > R {
			stack = append(stack, node.Left)
		} else {
			sum += node.Val
			stack = append(stack, node.Left, node.Right)
		}
	}
	return sum
}
