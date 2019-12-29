package main

import ."LeetCode-go/utils"

//中序遍历：左节点-根节点-右节点

//递归版本：Time：O(n), Space:(n)
func inorderTraversal(root *TreeNode) []int {
	if root == nil { //递归终止条件：root为空时，返回空list
		return []int{}
	}

	left := inorderTraversal(root.Left) //递归左子树
	right := inorderTraversal(root.Right) //递归右子树

	//最终的结果集【左子树的值，根节点的值，右子树的值】
	left = append(left, root.Val)
	left = append(left, right...) //注意此处将右子树的值加入结果集中的写法
	return left
}

//迭代版本：Time：O(n), Space:(n)
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	var res []int

	//注意判断的条件是或
	for root != nil || len(stack) != 0{
		//1.从根节点开始，将最左侧支路节点全部入栈
		//2.后续循环中，root代表栈顶元素的右节点，也就是右子树的根节点，入栈后继续优先将右子树的左节点入栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		//弹出栈顶元素
		length := len(stack) - 1
		root = stack[length]
		stack = stack[:length]
		//将栈顶元素的值加入结果集
		res = append(res, root.Val)
		//访问弹出元素的右子树
		root = root.Right
	}
	return res
}
