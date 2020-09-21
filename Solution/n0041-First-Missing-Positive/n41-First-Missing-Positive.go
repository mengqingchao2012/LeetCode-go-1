package main

func firstMissingPositive(nums []int) int {
	i, n := 0, len(nums)

	// Cyclist Sort 先进行原地排序
	for i < n {
		j := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[j] {
			v := nums[i] - 1
			nums[i], nums[v] = nums[v], nums[i]
		} else {
			i++
		}
	}

	// 然后再重头遍历进行比较
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
