package main

//Time：O(lgn), Space: O(1)
func countDigitOne(s int) int {
	if s < 1 {
		return 0
	}

	n := int64(s)
	var count, factor int64 = 0, 1 //count用于计数，factor是当前的位数，1是个位
	for n / factor != 0 {
		var digit = (n / factor) % 10
		var high = n / (10 * factor)
		switch digit {
		case 0:
			count += high * factor
		case 1:
			count += high * factor
			count += (n % factor) + 1
		default:
			count += (high + 1) * factor
		}
		factor *= 10
	}
	return int(count)
}
