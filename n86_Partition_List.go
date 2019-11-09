package main

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	smaller := &ListNode{}
	greater := &ListNode{}
	ps := smaller
	pg := greater

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			ps.Next = cur
			ps = ps.Next
		} else {
			pg.Next = cur
			pg = pg.Next
		}
	}
	ps.Next = greater.Next
	pg.Next = nil

	return smaller.Next
}
