package main

func nextPermutation(nums []int) {
	if nums == nil {
		return
	}

	n := len(nums)
	p := n - 2

	for p >= 0 && nums[p] >= nums[p+1] {
		p--
	}

	if p >= 0 {
		i := n - 1
		for i > p && nums[i] <= nums[p] {
			i--
		}
		nums[p], nums[i] = nums[i], nums[p]
	}

	j := p + 1
	for i := n - 1; j < i; {
		nums[j], nums[i] = nums[i], nums[j]
		j++
		i--
	}
}
