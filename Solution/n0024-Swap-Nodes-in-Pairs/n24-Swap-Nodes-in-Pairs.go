package main

import . "LeetCode-go/utils"

//迭代解法：Time：O(n), Space：O(1)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		start := pre.Next
		end := pre.Next.Next
		pre.Next = end
		start.Next = end.Next
		end.Next = start
		pre = start
	}
	return dummy.Next
}

//递归解法：Time：O(n), Space：O(1)
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}
