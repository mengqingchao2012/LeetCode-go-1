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
	length := len(nums)

	p := length - 2
	for p >= 0 && nums[p] >= nums[p+1] {
		p--
	}

	if p < 0 {
		return false
	}

	i := length - 1
	for i > p && nums[i] <= nums[p] {
		i--
	}
	nums[i], nums[p] = nums[p], nums[i]

	i, j := p+1, length-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return true
}
