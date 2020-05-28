package main

import . "LeetCode-go/utils"

//递归解法
func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	val := dfs(root.Right, sum)
	root.Val += val
	sum = root.Val
	return dfs(root.Left, sum)
}

//迭代解法
func convertBST1(root *TreeNode) *TreeNode {
	stack := []*TreeNode{}
	sum := 0
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur.Val += sum
		sum = cur.Val
		cur = cur.Left
	}
	return root
}
