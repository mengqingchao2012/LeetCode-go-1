package main

import (
	"container/heap"
)

//分治法：Time：O(nlogk), Space：O(logk), n是总节点数，k是链表的个数
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	}
	return merges(&lists, 0, length - 1)
}

func merges(lists *[]*ListNode, start, end int) *ListNode{
	if start == end {
		return (*lists)[start]
	}

	if start > end {
		return nil
	}

	mid := start + ((end - start) >> 1)
	left := merges(lists, start, mid)
	right := merges(lists, mid + 1, end)
	return mergeTwoSortedLists(left, right)
}

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
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

//方法二：构造小顶堆，Time：O(nlogk), Space：O(k), n是总节点数，k是链表的个数
type ListNode1 struct {
	Val int
	Next *ListNode1
}

func mergeKLists1(lists []*ListNode1) *ListNode1 {
	if len(lists) == 0 {
		return nil
	}

	temp := &minHeap{}
	heap.Init(temp)
	for _, list := range lists {
		if list != nil {
			heap.Push(temp, list)
		}
	}

	res := &ListNode1{}
	cur := res

	for temp.Len() != 0 {
		min := heap.Pop(temp).(*ListNode1)
		cur.Next = min
		cur = cur.Next
		if min.Next != nil {
			heap.Push(temp, min.Next)
		}
	}
	return res.Next
}

type minHeap []*ListNode1

func (mp minHeap) Len() int {
	return len(mp)
}

func (mp minHeap) Swap(i, j int) {
	mp[i], mp[j] = mp[j], mp[i]
}

func (mp minHeap) Less(i, j int) bool {
	return mp[i].Val < mp[j].Val
}

func (mp *minHeap) Push(h interface{}) {
	*mp = append(*mp, h.(*ListNode1))
}

func (mp *minHeap) Pop() (x interface{}) {
	n := len(*mp)
	res := (*mp)[n - 1]
	*mp = (*mp)[: n - 1]

	return res
}

//func main() {
//	a := &ListNode1{Val: 1}
//	b := &ListNode1{Val: 4}
//	c := &ListNode1{Val: 5}
//	a.Next = b
//	b.Next = c
//
//	d := &ListNode1{Val: 1}
//	e := &ListNode1{Val: 3}
//	f := &ListNode1{Val: 4}
//	d.Next = e
//	e.Next = f
//
//	g := &ListNode1{Val: 2}
//	h := &ListNode1{Val: 6}
//	g.Next = h
//
//	q := []*ListNode1{a, d, g}
//
//
//	//q := []*ListNode1{&ListNode1{Val:1}, &ListNode1{Val:0}}
//
//	res := mergeKLists(q)
//
//	for res != nil {
//		fmt.Println(res.Val)
//		res = res.Next
//	}
//}
