package main

//位运算：两个相同的数异或的结果为0，不相同的数异或结果为1
func singleOnlyNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}
