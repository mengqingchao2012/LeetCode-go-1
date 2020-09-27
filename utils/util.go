package utils

//链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//求两数的最大值
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//求两数之差（结果返回正数）
func Differ(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

//求两数的最小值
func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Abs64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func MultiMax(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > res {
			res = a[i]
		}
	}
	return res
}

func MultiMin(a ...int) int {
	res := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}