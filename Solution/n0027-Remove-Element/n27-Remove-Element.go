package main

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	slow, fast := 0, 0
	for ; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
