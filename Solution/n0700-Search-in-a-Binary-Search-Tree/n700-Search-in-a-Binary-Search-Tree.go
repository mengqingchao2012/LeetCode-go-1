package main

import . "LeetCode-go/utils"

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}
