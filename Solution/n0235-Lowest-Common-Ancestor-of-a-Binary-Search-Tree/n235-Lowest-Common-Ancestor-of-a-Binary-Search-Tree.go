package main

import ."LeetCode-go/utils"

//Time：O(h)，Space：O(h)，h是树高
//236题的特殊情况
//抓住题目特征：二叉搜索树，满足，左子树的值都小于根节点，右子树的值都大于根节点
func lowestCommonBinarySearchTreeAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val { //如果p,q都小于根节点，说明都在左子树
		return lowestCommonBinarySearchTreeAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val { //如果p,q都小于根节点，说明都在右子树
		return lowestCommonBinarySearchTreeAncestor(root.Right, p, q)
	} else { //p,q在根节点两边，则根节点就是最近的公共祖先
		return root
	}
}
