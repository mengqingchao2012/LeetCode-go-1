package n0026_Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[slow] == nums[fast] {
			continue
		}
		slow++
		nums[slow] = nums[fast]
	}
	return slow + 1
}
