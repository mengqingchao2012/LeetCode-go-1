package main

import ."LeetCode-go/utils"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n { // m == n，直接返回
		return head
	}
	// 注意：不要用 pre := &ListNode{} 的方式初始化，这样得到的 pre 不是 Nil
	var pre *ListNode
	pre, cur := nil, head

	// 注意下标的范围，以及不要忘记判断 cur != nil
	for i := 0; i < m - 1 && cur != nil; i++ {
		pre = cur
		cur = cur.Next
	}
	// 上面的循环结束后，pre 指向的是位置 m 的前一个节点，cur 指向位置 m 对应的节点
	lastNodeInFirstPart := pre // 表示节点位置 [:m) 的最后一个节点
	lastNodeInSubList := cur // 表示翻转前节点位置 [m:n] 的第一个节点，即 m，也是翻转后的最后一个节点

	// 翻转
	for i := 0; i < n - m + 1 && cur != nil; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 注意这里的条件判断：lastNodeInFirstPart == nil，意味着 m = 1，
	// 即从第一个元素开始翻转，故直接将 head 设置为翻转后翻转区间内的第一个节点
	if lastNodeInFirstPart != nil {
		lastNodeInFirstPart.Next = pre
	} else {
		head = pre
	}

	lastNodeInSubList.Next = cur
	return head
}
