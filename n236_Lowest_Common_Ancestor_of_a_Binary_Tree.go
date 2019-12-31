package main

import ."LeetCode-go/utils"

//方法一：
//思路：
//1.对某个节点，如果节点为空，或者节点等于 p, q 中的任一值，则直接返回该节点
//2.如果不满足条件1，则分别递归的遍历其左右子树，在左右子树上寻找 p，q
//3.如果某个节点的左右子树都有返回值，说明 p, q 各在一颗子树上，节点即是最近祖先
//4.如果某个节点只有一颗子树上有返回值，则说明有一个节点在该子树上，直接返回不为空的值
//Time：O(n)，Space：O(n)，n是节点总数
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

//方法二：
//思路：
//1.观察树结构可知，某个节点的所有祖先就等于从根节点到该节点路径上所有节点的集合（包括该节点本身）
//2.由此可以依次求出 p, q 两个节点的所有祖先，然后找到第一个不相同的节点，也就找到了最近的公共祖先

//该函数用于在二叉树中查找某个节点，同时记录下其路径上的所有节点
//函数返回一个bool类型的值，用来指示是否找到该节点
func searchTreeNode(root, node *TreeNode, path *[]*TreeNode) bool {
	if root == nil { //根节点为nil，直接返回false
		return false
	}
	*path = append(*path, root) //否则的话将该根节点加入到路径中
	if root == node { //如果根节点的值等于要查找的节点值，则直接返回
		return true
	}
	//否则，分别递归在该节点的左右子树上查找目标节点
	ret := searchTreeNode(root.Left, node, path) || searchTreeNode(root.Right, node, path)
	if ret { //只要有任意一棵子树查找到目标节点，说明当前节点就是目标节点的祖先之一
		return true
	}
	*path = (*path)[:len(*path) - 1] //否则的话说明目标节点不在当前节点的后代中，移除当前节点，返回false
	return false
}

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
	return ppath[i - 1]
}


