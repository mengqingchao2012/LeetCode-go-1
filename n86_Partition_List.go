package main

import "LeetCode-go/utils"

func partition(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil {
		return nil
	}

	smaller := &utils.ListNode{}
	greater := &utils.ListNode{}
	ps := smaller
	pg := greater

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			ps.Next = cur
			ps = ps.Next
		} else {
			pg.Next = cur
			pg = pg.Next
		}
	}
	ps.Next = greater.Next
	pg.Next = nil

	return smaller.Next
}
