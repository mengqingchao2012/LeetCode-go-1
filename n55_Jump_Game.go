package main

//Time：O(n); Space：O(1)
func canJump(nums []int) bool {
	size := len(nums)
	if size == 0 {
		return false
	}

	maxIndex := 0
	for i:= 0; i < size; i++ {
		if maxIndex >= size - 1 { //最大可达下标超过数组长度，说明可达
			return true
		}
		if i > maxIndex { //当前下标超过可达下标，说明不可达
			return false
		}
		maxIndex = max(maxIndex, i + nums[i]) //当前位置i处的最大可达下标为 i+nums[i]，与maxIndex比较更新maxIndex
	}
	return false
}
