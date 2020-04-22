package main

//Time：O(n*n!), Space:(n)
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	var result [][]int
	permuteRec(nums, 0, &result)
	return result
}

func permuteRec(nums []int, start int, res *[][]int) {
	length := len(nums)
	if start == length { //递归终止条件
		temp := make([]int, length)
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}

	for i := start; i < length; i++ {
		nums[i], nums[start] = nums[start], nums[i]
		permuteRec(nums, start+1, res)
		nums[i], nums[start] = nums[start], nums[i]
	}
}