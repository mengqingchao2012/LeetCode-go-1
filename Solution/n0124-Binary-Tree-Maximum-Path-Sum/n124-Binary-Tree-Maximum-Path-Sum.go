package main

import (
	. "LeetCode-go/utils"
	"math"
)

// 方法一：将每个节点都看成根节点，对其求最大路径和，最终再将结果进行对比
func maxPathSum(root *TreeNode) int {
	// 如果根节点为 null，直接返回 MinInt
	if root == nil {
		return math.MinInt32
	}
	// 最大路径和分为两种情况：经过根节点和不经过根节点
	// 路径和可能为负，负值路径和对结果没有正向贡献，因此可以进行过滤
	leftMax := Max(pathSumCalculate(root.Left), 0)
	rightMax := Max(pathSumCalculate(root.Right), 0)
	return Max(root.Val + leftMax + rightMax, Max(maxPathSum(root.Left), maxPathSum(root.Right)))
}

// 该方法通过传入一个节点，返回以该节点为根节点的最大路径和
func pathSumCalculate(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 递归求解左右子树的路径和，则该节点的最大路径和就是左右子树中的最大路径加上当前节点的值
	// 注意对负路径和进行过滤，如果左右子树的路径和为负，直接舍去只使用当前节点的值
	leftTreePathSum := Max(pathSumCalculate(root.Left), 0)
	rightTreePathSum := Max(pathSumCalculate(root.Right), 0)
	return root.Val + Max(leftTreePathSum, rightTreePathSum)
}

// 方法二：将求解最大路径和的步骤下推到遍历结点的过程中，一边遍历一边更新，这样只需要遍历一次节点即可
func maxPathSum1(root *TreeNode) int {
	res := math.MinInt32
	pathSum(root, &res)
	return res
}

func pathSum(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}

	leftTreePathSum := Max(pathSum(root.Left, res), 0)
	rightTreePathSum := Max(pathSum(root.Right, res), 0)
	*res = Max(*res, root.Val + leftTreePathSum + rightTreePathSum)
	return root.Val + Max(leftTreePathSum, rightTreePathSum)
}
