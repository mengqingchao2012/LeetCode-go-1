package main

import ."LeetCode-go/utils"

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

// 给定左右边界的值，生成在该区间内所有的二叉搜索树
func generate(low, high int) []*TreeNode {
	// 注意这里返回值一定要写成 []*TreeNode{nil}, 写成 []*TreeNode{} 会导致错误的结果
	// 退递归条件
	if low > high { return []*TreeNode{nil} }
	res := []*TreeNode{}

	for i := low; i <= high; i++ {
		left := generate(low, i - 1)
		right := generate(i + 1, high)
		// 以 i 为根节点，从其左子树的所有可能结果中选一颗作为左子树，同理选出一颗右子树
		// 创建二叉搜索树
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = copyTree(l)
				root.Right = copyTree(r)
				res = append(res, root)
			}
		}
	}
	return res
}

// 为了满足每颗子树都是独立的这个条件，手动拷贝每颗树
func copyTree(tree *TreeNode) *TreeNode {
	if tree == nil {
		return nil
	}
	root := &TreeNode{Val: tree.Val}
	root.Left = copyTree(tree.Left)
	root.Right = copyTree(tree.Right)
	return root
}