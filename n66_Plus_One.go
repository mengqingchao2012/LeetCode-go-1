package main

//Time：O(n)，Space:O(1)
func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry != 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
