package main

import ."LeetCode-go/utils"

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root, root)
}

func checkSymmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil { return true } // 注意判断的顺序，先判断 &&，再判断 ||，否则可能因为短路原则出错
	if t1 == nil || t2 == nil { return false }

	return t1.Val == t2.Val && checkSymmetric(t1.Left, t2.Right) && checkSymmetric(t1.Right, t2.Left)
}
