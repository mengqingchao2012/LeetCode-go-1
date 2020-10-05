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
	// 递归终止条件，start == length，说明已经取完了所有的元素，得到一组解
	if start == length {
		temp := make([]int, length)
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}

	// 遍历从 start 开始的子数组
	for i := start; i < length; i++ {
		// 将子数组中的元素依次与 start 进行交换，相当于固定不同的元素到子数组的第一个位置
		nums[i], nums[start] = nums[start], nums[i]
		// 递归求解从 start + 1 开始的子数组的全排列
		permuteRec(nums, start+1, res)
		// 退递归时要将 start 再换回原来的位置
		nums[i], nums[start] = nums[start], nums[i]
	}
}
