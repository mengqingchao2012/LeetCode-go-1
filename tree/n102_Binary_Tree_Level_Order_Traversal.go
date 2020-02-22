package main

import . "LeetCode-go/utils"

//递归解法
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}

//迭代解法：Time：O(n)，Space：O(n)
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	var res [][]int

	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		var elem []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			elem = append(elem, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, elem)
	}
	return res
}
