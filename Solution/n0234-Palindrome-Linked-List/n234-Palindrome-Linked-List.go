package n0234_Palindrome_Linked_List

import . "LeetCode-go/utils"

//Time：O(n)，Space:O(1)
func isPalindromeLinked(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil { //找到中点
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil { //如果链表长度为奇数，slow位于中间结点，需要对比的是中间结点的两侧结点，中间结点不对比
		slow = slow.Next //故将slow结点后移一位
	}

	slow = reversed(slow) //反转后半部分链表

	for slow != nil {
		if head.Val != slow.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func reversed(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
