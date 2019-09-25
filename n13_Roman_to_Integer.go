package main

// Time：O(n)，Space：O(1)
func romanToInt(s string) int {
	mapping := make([]int, 256)

	mapping['I'] = 1
	mapping['V'] = 5
	mapping['X'] = 10
	mapping['L'] = 50
	mapping['C'] = 100
	mapping['D'] = 500
	mapping['M'] = 1000

	//如果一个罗马数字从左向右递减，则转为阿拉伯数字只需将各个数字相加即可；
	//若果出现右边大于左边的情况，相当于减去左边的数得到
	size := len(s)
	res := mapping[s[size - 1]]
	for i := size - 2; i >= 0; i-- {
		cur := mapping[s[i]]
		right := mapping[s[i + 1]]
		if cur < right {
			res -= cur
		} else {
			res += cur
		}
	}
	return res
}
