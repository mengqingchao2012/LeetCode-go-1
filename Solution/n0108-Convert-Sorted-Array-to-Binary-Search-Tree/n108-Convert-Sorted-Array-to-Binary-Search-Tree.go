package main

import . "LeetCode-go/utils"

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return recursiveStruct(nums, 0, len(nums)-1)
}

func recursiveStruct(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + ((end - start) >> 1)
	treeRoot := &TreeNode{Val: nums[mid]}
	treeRoot.Left = recursiveStruct(nums, start, mid-1)
	treeRoot.Right = recursiveStruct(nums, mid+1, end)
	return treeRoot
}
