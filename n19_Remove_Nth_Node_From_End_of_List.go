package main

//Time：O(n)，Space:O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	p, q := dummy, dummy

	for ; n > 0 && q.Next != nil; n-- { //q指针先前移n步，注意边界条件
		q = q.Next
	}

	if n != 0 { //如果上一步循环退出后n不等于0，说明n大于链表长度，直接返回链表
		return dummy.Next
	}

	for q.Next != nil { //同时移动p, q，直到q移动到链表尾，此时p的next指向的即是倒数第n个数
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next
	return dummy.Next
}
