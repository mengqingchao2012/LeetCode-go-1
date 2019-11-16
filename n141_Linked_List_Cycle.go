package main

import "LeetCode-go/utils"

func hasCycle(head *utils.ListNode) bool {
	slow := head
	fast := head

	for fast!= nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
