package main

import ."LeetCode-go/utils"

type BSTIterator struct {
	root *TreeNode
	nodes []int
	size int
	idx int
}

// Time: O(n)，Space：O(n)
func Constructor(root *TreeNode) BSTIterator {
	var trival func (*TreeNode)
	nodes := []int{}

	trival = func (r *TreeNode) {
		if r == nil {
			return
		}

		trival(r.Left)
		nodes = append(nodes, r.Val)
		trival(r.Right)
	}
	trival(root)

	return BSTIterator {
		root: root,
		nodes: nodes,
		size: len(nodes),
	}
}

// Time: O(1)，Space：O(1)
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.HasNext() {
		res := this.nodes[this.idx]
		this.idx++
		return res
	}
	return -1
}


// Time: O(1)，Space：O(1)
/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.idx == this.size {
		return false
	}
	return true
}