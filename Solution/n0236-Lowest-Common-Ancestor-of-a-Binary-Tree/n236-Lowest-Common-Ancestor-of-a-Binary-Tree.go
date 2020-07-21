package main

import ."LeetCode-go/utils"

// 方法一
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	} else {
		return right
	}
}

//该函数用于在二叉树中查找某个节点，同时记录下其路径上的所有节点
//函数返回一个bool类型的值，用来指示是否找到该节点
func searchTreeNode(root, node *TreeNode, path *[]*TreeNode) bool {
	//根节点为nil，直接返回false
	if root == nil {
		return false
	}
	//否则的话将该根节点加入到路径中
	*path = append(*path, root)
	//如果根节点的值等于要查找的节点值，则直接返回
	if root == node {
		return true
	}
	//否则，分别递归在该节点的左右子树上查找目标节点,只要有任意一棵子树查找到目标节点，说明当前节点就是目标节点的祖先之一
	if searchTreeNode(root.Left, node, path) || searchTreeNode(root.Right, node, path) {
		return true
	}
	*path = (*path)[:len(*path)-1] //否则的话说明目标节点不在当前节点的后代中，移除当前节点，返回false
	return false
}

// 方法二
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	var ppath, qpath []*TreeNode

	searchTreeNode(root, p, &ppath)
	searchTreeNode(root, q, &qpath)

	//选择 p, q 两个节点的路径中较短的那一条，依次判断两条路径上相同下标对应的元素是否相等，找到第一个不相等的值时
	//其前一个元素就是最近公共祖先
	i := 0
	length := Min(len(ppath), len(qpath))
	for i < length && ppath[i] == qpath[i] {
		i++
	}
	return ppath[i-1]
}
