package main

import ."LeetCode-go/utils"

// 方法一：快排（不交换值）
// 思路：定义 left, mid, right 三个新的头结点，选择 head.val 作为分割点，遍历链表，
// 如果当前值等于分隔点的值，则将节点加入 mid 中，小于则加入 left，大于则加入 right，
// 然后递归对 left 和 right 进行排序，最后再将三部分拼接起来
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	left, mid, right := &ListNode{}, &ListNode{}, &ListNode{}
	// 三个指针来遍历三个链表，同时用来存储各自的尾结点
	ltail, midtail, rtail := left, mid, right
	pivot := head.Val

	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == pivot {
			midtail.Next = cur
			midtail = midtail.Next
		} else if cur.Val < pivot {
			ltail.Next = cur
			ltail = ltail.Next
		} else {
			rtail.Next = cur
			rtail = rtail.Next
		}
	}
	// 注意遍历完后要将尾结点的 next 指针设置为 nil
	ltail.Next, midtail.Next, rtail.Next = nil, nil, nil
	left.Next = sortList(left.Next)
	right.Next = sortList(right.Next)

	getTail(left).Next = mid.Next
	// 因为 mid.Next 的值可能为空，谨慎起见还是选择重头遍历来寻找尾结点
	getTail(left).Next = right.Next
	return left.Next
}

// 给定一个头结点，找到它的尾结点
func getTail(head *ListNode) *ListNode {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

// 方法二：快排（交换值）
func sortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}

	pivot := head.Val
	slow, fast := head, head.Next
	for fast != end {
		if fast.Val <= pivot {
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}
		fast = fast.Next
	}
	slow.Val, head.Val = head.Val, slow.Val

	quickSort(head, slow)
	quickSort(slow.Next, end)
}

// 方法三：归并排序
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	// 注意此处的条件判定！！！！！
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	right := sortList2(slow.Next)
	slow.Next = nil
	left := sortList2(head)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return res.Next
}
