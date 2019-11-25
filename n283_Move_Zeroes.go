package main

//Time：O(n)，Space：O(1)
func moveZeroes(nums []int) {
	length := len(nums)

	if length == 0 {
		return
	}

	slow, fast := 0, 0
	for ; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
