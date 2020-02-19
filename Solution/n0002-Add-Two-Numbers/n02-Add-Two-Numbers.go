package main

import "LeetCode-go/utils"

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	carry := 0
	res := &utils.ListNode{}
	cur := res

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		cur.Next = &utils.ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return res.Next
}

