package main

import (
	"container/heap"
	"math"
)

func smallestRange(nums [][]int) []int {
	hp := new(minheap)
	heap.Init(hp)

	minRange, maxRange, currentMax := math.MinInt32, math.MaxInt32, math.MinInt32
	for idx := range nums {
		value := nums[idx][0]
		heap.Push(hp, [3]int{idx, 0, value})
		if currentMax < value {
			currentMax = value
		}
	}

	size := len(nums)
	for hp.Len() == size {
		elem := heap.Pop(hp).([3]int)
		if maxRange - minRange > currentMax - elem[2] {
			minRange = elem[2]
			maxRange = currentMax
		}

		elem[1]++
		if elem[1] < len(nums[elem[0]]) {
			value := nums[elem[0]][elem[1]]
			elem[2] = value
			heap.Push(hp, elem)
			if currentMax < value {
				currentMax = value
			}
		}
	}
	return []int{minRange, maxRange}
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

