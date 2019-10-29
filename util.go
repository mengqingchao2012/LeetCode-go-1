package main

//链表
type ListNode struct {
	Val int
	Next *ListNode
}

//求两数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//求两数之差（结果返回正数）
func differ(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
//求两数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}