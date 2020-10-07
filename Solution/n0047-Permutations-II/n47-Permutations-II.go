package main

import "sort"

func permuteUnique(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return [][]int{}
	}

	res := [][]int{}
	sort.Ints(nums) // 为了执行 nextPermutation，要先将原数组按升序进行排序

	tmp := make([]int, length)
	copy(tmp, nums)
	res = append(res, tmp) // 注意这里要执行深拷贝，否则后续修改 nums 的值也会影响到最终结果

	for nextPermutation(nums) {
		// 同样注意深拷贝，不能复用之前定义的 tmp 变量，必须重新初始化
		tmp := make([]int, length)
		copy(tmp, nums)
		res = append(res, tmp)
	}

	return res
}

func nextPermutation(nums []int) bool {
	n := len(nums)
	p := n - 2
	for p >= 0 && nums[p] >= nums[p + 1] {
		p--
	}

	if p >= 0 {
		for i := n - 1; i > p; i-- {
			if nums[i] > nums[p] {
				nums[i], nums[p] = nums[p], nums[i]
				break
			}
		}
	}

	for i, j := p + 1, n - 1; i < j; i, j = i + 1, j - 1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return p != -1
}
