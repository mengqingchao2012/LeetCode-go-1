package main

import ."LeetCode-go/utils"

// 方法一：层序遍历（bfs）
func rightSideView(root *TreeNode) []int {
	if root == nil { return []int{} }
	queue := []*TreeNode{}
	res := []int{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		// 每切换到新的一层，就将该层的最右节点值加入到结果集中
		res = append(res, queue[0].Val)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 注意顺序是：根-右-左
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return res
}

// 方法二：dfs
func rightSideView1(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, res *[]int, level int) {
	if root == nil { return }
	if level == len(*res) { // 注意结果集要使用指针
		*res = append(*res, root.Val)
	}
	dfs(root.Right, res, level + 1)
	dfs(root.Left, res, level + 1)
}