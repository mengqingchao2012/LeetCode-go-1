package main

import (
	. "LeetCode-go/utils"
)

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int
	var res *ListNode //注意此处不能使用 res := &ListNode{}，因为这样初始化得到的是ListNode{0， nil}，引入了额外的节点

	//遍历两个链表，将结果存入数组
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	//倒序取出数组中的结果，并完成相加和链表的构建
	size1, size2 := len(s1) - 1, len(s2) - 1
	carry := 0
	for size1 >= 0 || size2 >= 0 || carry > 0 {
		sum := carry
		if size1 >= 0 {
			sum += s1[size1]
			size1--
		}
		if size2 >= 0 {
			sum += s2[size2]
			size2--
		}
		val := sum % 10
		carry = sum / 10
		pre := &ListNode{Val: val, Next: res} //构建新节点，新节点的Next指针指向之前的节点
		res = pre //更新res为新节点
	}
	return res
}

//func main() {
//	a := &ListNode{Val: 7}
//	b := &ListNode{Val: 2}
//	c := &ListNode{Val: 4}
//	d := &ListNode{Val: 3}
//	a.Next = b
//	b.Next = c
//	c.Next = d
//
//	e := &ListNode{Val: 5}
//	f := &ListNode{Val: 6}
//	g := &ListNode{Val: 4}
//	e.Next = f
//	f.Next = g
//
//	res := addTwoNumbersII(a, e)
//	fmt.Println(res)
//}