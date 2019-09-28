package main

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1)
	}
	return count
}

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