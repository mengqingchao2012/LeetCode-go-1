package main

import "sort"

type RecentCounter struct {
	queue []int
}

func Constructor1() RecentCounter {
	var q []int
	return RecentCounter{
		queue: q,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	boundary := t - 3000
	for this.queue[0] < boundary {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}

//解法二：使用二分法查找到 t-3000 的下标，然后用数组长度减去下标
func (this *RecentCounter) Ping1(t int) int {
	this.queue = append(this.queue, t)
	return len(this.queue) - sort.SearchInts(this.queue, t-3000)
}
