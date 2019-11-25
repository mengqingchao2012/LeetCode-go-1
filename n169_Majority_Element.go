package main

//方法一：摩尔投票法：Time：O(n)，Space：O(1)
//思路：众数只有一个，记下当前数，并初始化计数器为1，遍历数组，如果下一个数与当前数相同，则计数器加1，否则计数器减一，
//并更新当前数，如果计数器减到了0，则重新初始化计数器为1，并记录下当前数
func majorityElement(nums []int) int {
	size := len(nums)
	num, count := nums[0], 1 //两个计数器，一个记录当前数，
	for i := 1; i < size; i++ {
		if count == 0 {
			num = nums[i]
			count++
		} else if nums[i] == num {
			count++
		} else {
			count--
		}
	}
	return num
}

//方法二：Hash表：Time：O(n)，Space：O(n)
func majorityElement1(nums []int) int {
	max := 0
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num] += 1
		if mp[num] > max {
			max = mp[num]
		}
		if max > len(nums)/2 {
			return num
		}
	}

	return -1
}

//func main() {
//	res := majorityElement([]int{3, 1, 3})
//	fmt.Println(res)
//}
