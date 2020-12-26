package main

type TreeArry struct {
	tr []int
	n int
}

func (t *TreeArry) lowBit(x int) int {
	// x & (~x + 1)
	return x & (-x)
}

// Time: O(logn)
func (t *TreeArry) Query(x int) int {
	res := 0
	for i := x; i > 0; i -= t.lowBit(i) {
		res += t.tr[i]
	}
	return res
}

// Time: O(logn)
func (t *TreeArry) Add(x int, value int) {
	for i := x; i <= t.n; i += t.lowBit(i) {
		t.tr[i] += value
	}
}

type NumArray struct {
	nums []int
	trrArr TreeArry
}

// Time: O(nlogn)
func Constructor(nums []int) NumArray {
	size := len(nums)
	trrArr := TreeArry {
		tr: make([]int, size + 1),
		n: size,
	}
	for i := 0; i < size; i++ {
		trrArr.Add(i + 1, nums[i])
	}

	return NumArray {
		nums: nums,
		trrArr: trrArr,
	}
}

func (n *NumArray) Update(i int, val int)  {
	n.trrArr.Add(i + 1, val - n.nums[i])
	n.nums[i] = val
}


func (n *NumArray) SumRange(i int, j int) int {
	return n.trrArr.Query(j + 1) - n.trrArr.Query(i)
}
