package main

func removeDuplicates(nums []int) int {
	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j], nums[i] = nums[i], nums[j]
	}
	return j + 1
}
