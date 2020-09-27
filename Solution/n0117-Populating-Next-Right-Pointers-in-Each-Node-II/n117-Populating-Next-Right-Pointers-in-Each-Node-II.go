package main

import ."LeetCode-go/utils"

// 方法一：层序遍历，时间复杂度：O(n)，空间复杂度：O(n)
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// 这种方法直接修改了 root。node 是弹出的队首元素，queue[0] 是下一个
			// 要被弹出的元素，因此 node.Next 应该指向 queue[0]
			// 限制 i < size - 1 的原因是防止访问到下一层
			if i < size - 1 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

// 方法二：时间复杂度：O(n)，空间复杂度：O(1)
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	cur := root
	for cur != nil {
		// 在遍历当前层的同时为下一层设置 next 指针
		// 为了统一处理，采用 dummy 节点的思路
		dummy := &Node{}
		pre := dummy
		for cur != nil { // 遍历当前层的节点
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}
			cur = cur.Next // 切换到当前层的下一个节点
		}
		// 切换到下一层的第一个节点
		cur = dummy.Next
	}
	return root
}
