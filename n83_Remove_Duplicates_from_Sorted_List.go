package main

import "LeetCode-go/utils"

//Time：O(n)，Space：O(1)
func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	cur := head
	next := head.Next

	for next != nil {
		if cur.Val == next.Val {
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
		next = next.Next
	}
	return head
}
