package main

func sortColors(nums []int)  {
	low, high := 0, len(nums) - 1
	i := 0
	for i <= high {
		if nums[i] == 0 {
			nums[low], nums[i] = nums[i], nums[low]
			i++
			low++
		} else if nums[i] == 1 {
			i++
		} else {
			// 注意此处：nums[high] 和 nums[i] 交换后，nums[i] 的值可能是 0, 1, 2 中的任意结果，
			// 因此需要进一步比较，不能跳过
			nums[high], nums[i] = nums[i], nums[high]
			high--
		}
	}
}
