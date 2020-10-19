package main

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	hp := new(maxheap)
	heap.Init(hp)

	for i, n1 := range nums1 {
		for j, n2 := range nums2 {
			if hp.Len() < k {
				heap.Push(hp, [3]int{i, j, n1 + n2})
			} else {
				if (*hp)[0][2] < n1 + n2 {
					break
				} else {
					heap.Pop(hp)
					heap.Push(hp, [3]int{i, j, n1 + n2})
				}
			}
		}
	}
	res := make([][]int, 0, k)
	for _, v := range *hp {
		res = append(res, []int{nums1[v[0]], nums2[v[1]]})
	}
	return res
}

type maxheap [][3]int

func (m maxheap) Len() int {
	return len(m)
}

func (m maxheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxheap) Less(i, j int) bool {
	return m[i][2] > m[j][2]
}

func (m *maxheap) Push(x interface{}) {
	*m = append(*m, x.([3]int))
}

func (m *maxheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[:n - 1]
	return res
}