package main

import (
	"container/heap"
	"math"
)

func kClosest(points [][]int, K int) [][]int {
	hp := new(minheap)
	heap.Init(hp)

	for _, v := range points {
		if hp.Len() < K {
			heap.Push(hp, v)
		} else if distance(v) < hp.Top() {
			heap.Pop(hp)
			heap.Push(hp, v)
		}
	}
	return *hp
}

type minheap [][]int

func (m minheap) Len() int {
	return len(m)
}

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minheap) Less(i, j int) bool {
	return distance(m[i]) > distance(m[j])
}

func (m *minheap) Top() float64 {
	return distance((*m)[0])
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.([]int))
}

func (m *minheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[:n - 1]
	return res
}

func distance(m []int) float64 {
	return math.Sqrt(float64(m[0] * m[0] + m[1] * m[1]))
}