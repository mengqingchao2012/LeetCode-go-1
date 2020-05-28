package main

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + carry
		digits[i] = tmp % 10
		carry = tmp / 10
	}
	if carry != 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 优化版
func plusOne1(digits []int) []int {
	p := len(digits) - 1           // 创建一个游标 p，指向数组的最后一个元素
	for p >= 0 && digits[p] == 9 { // 如果游标 p 大于等于0且游标 p 所指元素为9，则将该元素修正为0，p前移一位
		digits[p] = 0
		p--
	}

	if p == -1 {
		digits = append([]int{1}, digits...)
	} else {
		digits[p] += 1
	}
	return digits
}
