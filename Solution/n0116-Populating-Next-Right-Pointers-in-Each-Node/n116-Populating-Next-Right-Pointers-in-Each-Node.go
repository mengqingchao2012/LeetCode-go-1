package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 方法一：递归法
func connect(root *Node) *Node {
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

// 方法二：迭代法
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	left, cur := root, &Node{}

	for left.Left != nil {
		cur = left
		for cur != nil {
			cur.Left.Next = cur.Right
			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}
			cur = cur.Next
		}
		left = left.Left
	}
	return root
}
