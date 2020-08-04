package main

import ."LeetCode-go/utils"

type BSTIterator1 struct {
	stack []*TreeNode
}

func (this *BSTIterator1) pushLeft(root *TreeNode) {
	for root != nil {
		this.stack = append(this.stack, root)
		root = root.Left
	}
}

// Time: O(h)，Space：O(h)
func Constructor1(root *TreeNode) BSTIterator1 {
	var iter BSTIterator1
	iter.stack = []*TreeNode{}
	iter.pushLeft(root)
	return iter
}

// Time_avg: O(1)，Space_avg：O(h)
/** @return the next smallest number */
func (this *BSTIterator1) Next() int {
	if this.HasNext() {
		node := this.stack[len(this.stack) - 1]
		this.stack = this.stack[:len(this.stack) - 1]
		this.pushLeft(node.Right)
		return node.Val
	}
	return -1
}

// Time_avg: O(1)，Space_avg：O(h)
/** @return whether we have a next smallest number */
func (this *BSTIterator1) HasNext() bool {
	return len(this.stack) != 0
}