package main

import ."LeetCode-go/utils"

// 迭代
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}

	size := len(res)
	for i := 0; i < size / 2; i++ {
		res[i], res[size - i - 1] = res[size - i - 1], res[i]
	}
	return res
}

// 递归
func levelOrderBottom1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
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

	n := len(res)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return res
}
