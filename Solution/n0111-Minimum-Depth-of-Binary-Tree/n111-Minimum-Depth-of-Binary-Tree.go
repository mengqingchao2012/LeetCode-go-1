package main

import ."LeetCode-go/utils"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	cnt := 1
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return cnt
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		cnt++
	}
	return cnt
}
