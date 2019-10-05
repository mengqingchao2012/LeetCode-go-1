package main

//Time: O(n); Space: O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	oldTail := head
	var n = 1
	for ; oldTail.Next != nil; n++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	newTail := head
	for i := 0; i < n - k % n - 1; i++ {
		newTail = newTail.Next
	}
	res := newTail.Next
	newTail.Next = nil

	return res
}
