package main

//原地翻转：Time: O(n); Space: O(1)
func rotate(nums []int, k int) {
	length := len(nums)
	if length == 0 {
		return
	}

	idx := k % length //k的大小可能大于length，所以要去除成环的影响

	//注意翻转顺序和start, end的下标
	rotation(&nums, 0, length-1)   //先翻转整个数组
	rotation(&nums, 0, idx-1)      //再翻转前idx个数
	rotation(&nums, idx, length-1) //最后翻转后续的数
}

func rotation(nums *[]int, start, end int) {
	for start < end {
		(*nums)[start], (*nums)[end] = (*nums)[end], (*nums)[start]
		start++
		end--
	}
}
