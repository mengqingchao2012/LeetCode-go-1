package main

import (
	. "LeetCode-go/utils"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	dfs(root, "", &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*res = append(*res, path + strconv.Itoa(root.Val))
	} else {
		path += strconv.Itoa(root.Val) + "->"
		dfs(root.Left, path, res)
		dfs(root.Right, path, res)
	}
}
