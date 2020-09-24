package main

import "LeetCode-go/utils"

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	oldTail := head
	var n = 1 //注意计数取值为1
	for ; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	newTail := head
	// 注意位置的确定：n - k%n - 1，得到的是翻转后链表的尾结点下标
	for i := 0; i < n-k%n-1; i++ {
		newTail = newTail.Next
	}
	// 尾结点的下一个节点就是链表头
	res := newTail.Next
	newTail.Next = nil

	return res
}
