package main

//插入排序：最坏时间复杂度 = 平均时间复杂度 = O(n^2), 空间复杂度：O(1)
func insertSort(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	for i := 1; i < length; i++ {
		cur := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > cur {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = cur
	}
}
