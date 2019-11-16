package main

import "LeetCode-go/utils"

func deleteDuplicates2(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	dummy := &utils.ListNode{}
	dummy.Next = head

	pre := dummy
	cur := pre.Next
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}

		if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = pre.Next
		}
		cur = pre.Next
	}
	return dummy.Next
}
