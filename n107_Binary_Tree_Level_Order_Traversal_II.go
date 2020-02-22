package main

import . "LeetCode-go/utils"

//相似题：n102——二叉树的层次遍历

//递归解法
//解法其实就是在102题的基础上，求出层次遍历的结果后，将结果集逆序即可，
//逆序的过程可以选择遍历一半结果集，然后首尾对换
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)

	n := len(res)
	for i := 0; i < n/2; i++ {
		j := n - 1 - i
		res[i], res[j] = res[j], res[i]
	}
	return res
}

//写法二：选择在插入结果集时直接构造结果集，不用额外遍历，然而这种写法其实时间复杂度会更高
func levelOrderBottom1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			//关键行：在level比当前层高时，说明遍历到了下一层，因为逆序输出，所以下一层的结果应该在
			//当前层之前，等同于在当前层的前面插入了一个新行
			//时间复杂度其实会更高，因为每次都要遍历一次结果集
			res = append([][]int{[]int{}}, res...)
		}
		n := len(res)
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return res
}
