package main

import "LeetCode-go/utils"

//Time: O(n); Space: O(1)
func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	oldTail := head
	var n = 1 //注意计数取值为1
	for ; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	newTail := head
	for i := 0; i < n-k%n-1; i++ { //注意位置的确定：n - k%n - 1
		newTail = newTail.Next
	}
	res := newTail.Next
	newTail.Next = nil

	return res
}
