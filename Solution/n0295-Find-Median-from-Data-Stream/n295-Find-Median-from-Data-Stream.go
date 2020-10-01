package main

import (
	"container/heap"
	"math"
)

type MedianFinder struct {
	Right *minHeap
	Left *maxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	right := new(minHeap)
	heap.Init(right)
	left := new(maxHeap)
	heap.Init(left)
	return MedianFinder {
		Right: right,
		Left: left,
	}
}

func (this *MedianFinder) AddNum(num int)  {
	// 当左半部分元素个数为 0 或者待插入元素小于左边最大堆的堆顶元素时，说明应该插入到左边
	if this.Left.Len() == 0 || num <= this.Left.Top() {
		heap.Push(this.Left, num)
	} else {
		heap.Push(this.Right, num)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	this.rebalance()
	size := this.Left.Len() + this.Right.Len()
	if size & 1 == 1 {
		return float64(this.Left.Top())
	} else {
		return float64(this.Left.Top() + this.Right.Top())/2
	}
}

func (this *MedianFinder) rebalance() {
	// 约定左边元素大于等于右边元素，所以如果右边更多，则每次弹出堆顶元素加入左边
	for this.Left.Len() < this.Right.Len() {
		heap.Push(this.Left, heap.Pop(this.Right))
	}

	// 注意如果数组中的总元素个数为偶数个时，要求左右堆中元素相等，因此若此时左边
	// 堆中元素过多，要弹出堆顶元素加入到右边堆中
	if (this.Left.Len() + this.Right.Len()) & 1 == 0 {
		for this.Left.Len() > this.Right.Len() {
			heap.Push(this.Right, heap.Pop(this.Left))
		}
	}
}

type minHeap struct {
	heapArray
}

func (h minHeap) Less(i, j int) bool {
	return h.heapArray[i] < h.heapArray[j]
}

type maxHeap struct {
	heapArray
}

func (h maxHeap) Less(i, j int) bool {
	return h.heapArray[i] > h.heapArray[j]
}

type heapArray []int

func (h heapArray) Len() int {
	return len(h)
}

func (h heapArray) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapArray) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *heapArray) Pop() interface{} {
	n := len(*h)
	res := (*h)[n-1]
	*h = (*h)[:n-1]
	return res
}

func (h *heapArray) Top() int {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return math.MaxInt32
}
