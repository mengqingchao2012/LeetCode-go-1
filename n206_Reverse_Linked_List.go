package main

import (
	"LeetCode-go/utils"
)

func reverseList(head *utils.ListNode) *utils.ListNode {
	var pre *utils.ListNode
	pre, cur := nil, head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
