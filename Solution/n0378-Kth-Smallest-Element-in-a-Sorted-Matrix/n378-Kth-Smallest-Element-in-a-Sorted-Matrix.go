package main

import (
	. "LeetCode-go/utils"
	"container/heap"
)

// 方法一：最小堆
func kthSmallest(matrix [][]int, k int) int {
	hp := new(minheap)
	heap.Init(hp)
	// 由于数组满足在行列两个方向上递增，故第 k 小的数必然出现在前 k 行中
	minRow := Min(len(matrix), k)
	// 将前 k 行的第一个元素依次加入最小堆中，堆中元素的组织形式为：【行号，列号，值】
	for i := 0; i < minRow; i++ {
		heap.Push(hp, [3]int{i, 0, matrix[i][0]})
	}

	cols := len(matrix[0])
	// 执行 k - 1 次弹出操作，则剩余的堆顶元素就是第 k 小的数
	for i := 0; i < k - 1; i++ {
		elem := heap.Pop(hp).([3]int)
		// 弹出每个元素后，如果该元素所在的行中后续还有元素，则将后续元素加入堆中
		elem[1]++ // 注意这里要先进行递增操作，再判断是否超过列数限制
		if elem[1] < cols {
			elem[2] = matrix[elem[0]][elem[1]]
			heap.Push(hp, elem)
		}
	}
	return (*hp)[0][2]
}

type minheap [][3]int

func (m minheap) Len() int {
	return len(m)
}

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minheap) Less(i, j int) bool {
	return m[i][2] < m[j][2]
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.([3]int))
}

func (m *minheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[: n - 1]
	return res
}

// 方法二： 二分法
func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix) - 1
	left, right := matrix[0][0], matrix[n][n]
	for left < right {
		mid := left + ((right - left) >> 1)
		if nLessMids(matrix, n, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func nLessMids(matrix [][]int, row, mid int) int {
	count := 0
	for i := 0; i <= row; i++ {
		if matrix[i][row] <= mid {
			count += row + 1
		} else {
			for j := 0; j <= row && matrix[i][j] <= mid; j++ {
				count++
			}
		}
	}
	return count
}