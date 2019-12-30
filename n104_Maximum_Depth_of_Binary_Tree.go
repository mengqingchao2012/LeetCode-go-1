package main

import ."LeetCode-go/utils"

//递归法：Time：O(n)，Space：O(n)
func maxDepth(root *TreeNode) int {
	if root == nil { //如果root为空，返回深度为0
		return 0
	}

	//递归返回左右结点中深度最大的那个，注意最终结果要加1
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//迭代法：Time：O(n)，Space：O(n)
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode //使用一个辅助队列
	depth := 0
	queue = append(queue, root)

	//层序遍历的思路，遍历完一层，深度+1
	for len(queue) != 0 {
		size := len(queue) //记录当前层的元素个数
		for i := 0; i < size; i++ {
			//出队
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
