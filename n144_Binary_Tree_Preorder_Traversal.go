package main

import . "LeetCode-go/utils"

//前序遍历：根-左节点-右节点

//递归解法：Time：O(n), Space:(n)
func recursive(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	recursive(root.Left, res)
	recursive(root.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	recursive(root, &res)
	return res
}

//迭代解法：Time：O(n), Space:(n)
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	var stack []*TreeNode       //使用栈模拟递归
	stack = append(stack, root) //根节点先入栈

	for len(stack) != 0 { //当栈内元素不为空时
		//弹出栈顶元素
		length := len(stack) - 1
		node := stack[length]
		stack = stack[:length]
		//将栈顶元素的值加入结果集
		res = append(res, node.Val)

		//根据栈后进先出的原则，先加入右节点，再加入左节点
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
