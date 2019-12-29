package main

import ."LeetCode-go/utils"

//后序遍历：左节点-右节点-根节点

//递归版本：Time：O(n), Space:(n)
func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	postorder(root, &res)
	return res
}

//迭代版本：Time：O(n), Space:(n)
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var (
		pre *TreeNode //pre用来记录上一个出栈元素的值，避免重复访问
		stack []*TreeNode
		res []int
	)

	//遍历过程：
	for root != nil || len(stack) != 0 {
		if root != nil { //1.如果当前节点不为空，将节点入栈并访问其左子树，如果左子树不为空，继续入栈
			stack = append(stack, root)
			root = root.Left
		} else { //2.如果当前节点为空，检查栈顶元素是否有右子树，如有则访问右子树
			length := len(stack) - 1
			root = stack[length].Right
			//3.当栈顶元素的右子树为空，或右子树等于上一个出栈的节点元素，
			//说明栈顶元素的右子树已经完成遍历，可以出栈栈顶元素，并将其值写入结果集中
			if root == nil || root == pre {
				pre = stack[length]
				stack = stack[:length]
				res = append(res, pre.Val)
				root = nil
			}
		}
	}
	return res
}
