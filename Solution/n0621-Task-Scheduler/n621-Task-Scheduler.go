package main

import (
	. "LeetCode-go/utils"
	"container/heap"
)

// 方法一：最大堆
func leastInterval(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, v := range tasks {
		freq[v - 'A']++
	}

	hp := new(maxheap)
	heap.Init(hp)
	for _, v := range freq {
		if v != 0 {
			heap.Push(hp, v)
		}
	}

	res, idle := 0, 0
	for hp.Len() != 0 {
		res += n + 1
		idle = n + 1 - hp.Len()
		size := Min(n + 1, hp.Len())
		for i := 0; i < size; i++ {
			freq[i] = heap.Pop(hp).(int) - 1
		}
		for i := 0; i < size; i++ {
			if freq[i] != 0 {
				heap.Push(hp, freq[i])
			}
		}
	}
	return res - idle
}

type maxheap []int

func (m maxheap) Len() int {
	return len(m)
}

func (m maxheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxheap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *maxheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxheap) Pop() interface{} {
	n := len(*m)
	res := (*m)[n - 1]
	*m = (*m)[:n - 1]
	return res
}

// 方法二：公式计算
func leastInterval1(tasks []byte, n int) int {
	freq := make([]int, 26)
	for _, v := range tasks {
		freq[v - 'A']++
	}

	maxFreq, cnt := 0, 0
	for _, v := range freq {
		if v > maxFreq {
			maxFreq = v
			cnt = 1
		} else if v == maxFreq {
			cnt++
		}
	}

	res := (n + 1) * (maxFreq - 1) + cnt
	return Max(res, len(tasks))
}
