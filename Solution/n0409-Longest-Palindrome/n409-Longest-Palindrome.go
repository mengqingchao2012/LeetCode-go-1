package main

func longestPalindrome(s string) int {
	mp := map[byte]int{}
	for _, v := range []byte(s) {
		mp[v]++
	}

	flag := false
	count := 0
	for _, v := range mp {
		if v%2 == 0 {
			count += v
		} else { // 注意如果字符出现次数为奇数，则要减去1，取出其中的偶数部分来构成最终结果
			flag = true
			count += v - 1
		}
	}

	if flag { // 如果有字符数为奇数次，则可以取任意一个奇数字符放在中间，满足中心对称的条件
		return count + 1
	} else {
		return count
	}
}
