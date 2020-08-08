package main

import ."LeetCode-go/utils"

// 前序遍历：根 —— 左 —— 右
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	nodes := []*TreeNode{}
	preOrder(root, &nodes)

	cur := root
	for i := 1; i < len(nodes); i++ {
		cur.Left = nil
		cur.Right = nodes[i]
		cur = cur.Right
	}
}

func preOrder(root *TreeNode, res *[]*TreeNode) {
	if root == nil {
		return
	}

	*res = append(*res, root)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

// 自底向上：右 —— 左 —— 根
func flatten1(root *TreeNode)  {
	if root == nil {
		return
	}
	// 注意：这里不要用 prev := &TreeNode{} 这种声明方式，
	// 因为会默认构造出一个值为 0，左右子树为 nil 的节点
	var prev *TreeNode
	var helper func(*TreeNode)

	// TODO：有没有不使用闭包的实现方式
	helper = func (root *TreeNode) {
		if root == nil {
			return
		}

		helper(root.Right)
		helper(root.Left)
		root.Left = nil
		root.Right = prev
		prev = root
	}
	helper(root)
}
