package main

//Time：O(n)，Space:O(1)
//由题意可知，数组中只有0，1，2三个元素，则使用两个指针分别指向头尾，头尾放的是0和2，再使用一个指针遍历
//数组，将不是中间数1的元素都调换到头尾去
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	pre, cur, last := 0, 0, len(nums)-1
	for cur <= last {
		if nums[cur] > 1 {
			nums[cur], nums[last] = nums[last], nums[cur]
			last--
		} else if nums[cur] < 1 {
			nums[cur], nums[pre] = nums[pre], nums[cur]
			cur++
			pre++
		} else {
			cur++
		}
	}
	return
}
