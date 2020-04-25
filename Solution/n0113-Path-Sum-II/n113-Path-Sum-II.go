package main

import ."LeetCode-go/utils"

//递归法：Time：O(n)，Space：O(n)
func pathSum(root *TreeNode, sum int) [][]int {
	var (
		res  [][]int //用来存储结果集
		elem []int   //用来存储路径中的元素
	)
	pathFinder(root, sum, &elem, &res)
	return res
}

func pathFinder(root *TreeNode, sum int, elem *[]int, res *[][]int) {
	if root == nil {
		return
	}

	*elem = append(*elem, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum { //到达叶节点并且叶节点的值等于nil
		//将当前elem加入结果集，说明找到一条满足条件的路径
		temp := make([]int, len(*elem))
		copy(temp, *elem)
		*res = append(*res, temp)
		// 注意这里找到正确值后不能直接 return !!! 因为要找到所有符合的路径，所以要执行最后一步的退递归
	}
	//递归左右子树
	pathFinder(root.Left, sum-root.Val, elem, res)
	pathFinder(root.Right, sum-root.Val, elem, res)

	*elem = (*elem)[:len(*elem)-1]
}
