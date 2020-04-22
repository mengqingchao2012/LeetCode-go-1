package main

func nextPermutation(nums []int)  {
	length := len(nums)
	if length < 2 { // 元素个数小于2，直接返回
		return
	}

	p := length - 2
	// 从右往左遍历，找到非降序排列的第一个数
	for p >= 0 && nums[p] >= nums[p + 1] { p-- }

	// 如果存在这样一个数p，则从右往左开始，找到第一个比该数大的数i，交换p和i
	if p >= 0 {
		i := length - 1
		for i > p && nums[i] <= nums[p] { i-- }
		nums[i], nums[p] = nums[p], nums[i]
	}

	// 将p之后的数字整理成升序排列
	// 这一步实际上包含了原始数组是降序排列的情况
	for i, j := p + 1, length - 1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
