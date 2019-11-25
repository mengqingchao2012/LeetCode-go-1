package main

//Time：O(n)，Space:O(1)
func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1

	for left < right {
		temp := 0
		//面积等于两根边界中较小的那个乘以边界之间的距离，面积要最大，高度的影响更大，移动高度小的边界
		if height[left] < height[right] {
			temp = height[left] * (right - left)
			left++
		} else {
			temp = height[right] * (right - left)
			right--
		}

		if temp > res { //比较当前面积与最大面积，更新最大面积
			res = temp
		}
	}
	return res
}
