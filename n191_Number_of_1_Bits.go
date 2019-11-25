package main

//Time：O(k)，k是二进制中1的个数，Space：O(1)
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1) // 将num的二进制中最低位的1消除
	}
	return count
}

//Time：O(n)，n是二进制的位数，Space：O(1)
func hammingWeight1(num uint32) int {
	var mask uint32 = 1
	count := 0
	for mask != 0 {
		if (mask & num) != 0 {
			count++
		}
		mask <<= 1
	}
	return count
}
