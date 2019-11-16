package main

import "LeetCode-go/utils"

func detectCycle(head *utils.ListNode) *utils.ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			for ptr := head; ptr != slow; ptr = ptr.Next {
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
