package main

import ."LeetCode-go/utils"

// 结合第92题来实现
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1{
		return head
	}

	// 先遍历一遍数组统计节点个数，由此可以计算出没 k 个元素翻转一次，需要翻转多少次
	pre := head
	count := 0
	for pre != nil {
		count += 1
		pre = pre.Next
	}
	count /= k

	// 然后就是每次翻转 k 个元素，从 [m : m + k - 1]
	// 这里的 m 从 1 开始算起
	m := 1
	for i := 0; i < count; i++ {
		head = reverseFromKTON(head, m, m + k - 1)
		m += k
	}
	return head
}

func reverseFromKTON(head *ListNode, m, n int) *ListNode {
	if m == n {
		return head
	}

	var pre *ListNode
	pre, cur := nil, head
	for i := 0; i < m - 1 && cur != nil; i++ {
		pre = cur
		cur = cur.Next
	}

	lastNodeInFirst := pre
	lastNodeInSublist := cur
	for i := 0; i < n - m + 1 && cur != nil; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if lastNodeInFirst != nil {
		lastNodeInFirst.Next = pre
	} else {
		head = pre
	}

	lastNodeInSublist.Next = cur
	return head
}
