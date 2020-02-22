package main

import . "LeetCode-go/utils"

//递归求解：Time：O(n)，Space：O(n)
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	switch {
	case t1 == nil && t2 == nil:
		return nil
	case t1 == nil && t2 != nil:
		return t2
	case t1 != nil && t2 == nil:
		return t1
	default:
		t1.Val += t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
	}
	return t1
}
