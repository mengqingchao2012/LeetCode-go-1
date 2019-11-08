package main

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
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
