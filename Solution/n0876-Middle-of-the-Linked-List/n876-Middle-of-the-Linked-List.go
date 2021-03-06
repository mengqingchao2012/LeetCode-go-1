package n0876_Middle_of_the_Linked_List

import (
	"LeetCode-go/utils"
)

func middleNode(head *utils.ListNode) *utils.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
