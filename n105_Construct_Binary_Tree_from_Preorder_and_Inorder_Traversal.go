package main

import "LeetCode-go/utils"

func buildTree(preorder []int, inorder []int) *utils.TreeNode {
	size := len(preorder)
	inPos := make(map[int]int, size)
	for k, v := range inorder {
		inPos[v] = k
	}
	return build(&preorder, 0, size-1, 0, &inPos)
}

func build(pre *[]int, preStart, preEnd int, inStart int, inPos *map[int]int) *utils.TreeNode {
	if preStart > preEnd {
		return nil
	}

	root := &utils.TreeNode{Val: (*pre)[preStart]}
	rootIndex := (*inPos)[(*pre)[preStart]]
	leftLen := rootIndex - inStart

	root.Left = build(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = build(pre, preStart+leftLen+1, preEnd, rootIndex+1, inPos)
	return root
}

//func main() {
//	pre := []int{3, 9, 20, 15, 7}
//	in := []int{9, 3, 15, 20, 7}
//	buildTree(pre, in)
//}
