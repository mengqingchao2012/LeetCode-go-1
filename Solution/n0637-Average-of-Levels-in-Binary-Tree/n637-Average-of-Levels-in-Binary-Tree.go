package main

import ."LeetCode-go/utils"

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	res := []float64{}
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		tmp := float64(sum) / float64(size)
		res = append(res, tmp)
	}
	return res
}
