package main

import . "LeetCode-go/utils"

//Time：O(n^2), Space：O(1),
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	pre := dummy //pre指针指向有序区的开始元素，所以比较的是cur.Val和pre.Next.Val的大小
	cur := head //cur指针遍历链表，依次取出每个结点进行插入
	for cur != nil {
		next := cur.Next //指向cur指针的next结点，保证链表不会中断
		//优化：如果当前cur指针所指元素的值大于等于pre.Next所指元素的值，则无需重置pre，继续移动pre指针即可
		if pre.Next != nil && cur.Val < pre.Next.Val {
			pre = dummy
		}
		for pre.Next != nil && cur.Val >= pre.Next.Val {
			pre = pre.Next
		}
		cur.Next = pre.Next
		pre.Next = cur
		cur = next
	}
	return dummy.Next
}
