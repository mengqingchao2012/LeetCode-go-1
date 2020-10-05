package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	n := len(nums)
	if n == 0 {
		return [][]int{}
	}

	cur := make([]int, 0, n)
	res := [][]int{}
	// 先对数组进行排序
	sort.Ints(nums)
	backTrace(nums, 0, cur, &res)
	return res
}

func backTrace(nums []int, start int, cur []int, res *[][]int) {
	*res = append(*res, cur)
	for i := start; i < len(nums); i++ {
		// 注意这里是 i > start 而不是 i > 0，保证回溯时跳过重复元素就是保证同一个位置上
		// 不能取到同一个元素，而这个位置是由 start 来决定的
		if i > start && nums[i - 1] == nums[i] {
			continue
		}
		cur = append(cur, nums[i])
		backTrace(nums, i + 1, cur, res)
		cur = cur[:len(cur) - 1]
		cur = cur[:len(cur):len(cur)]
	}
}