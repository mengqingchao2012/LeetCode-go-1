package main

func countCharacters(words []string, chars string) int {
	charSlice := make([]int, 26)
	for _, v := range []byte(chars) {
		charSlice[v - 'a']++
	}

	temp := make([]int, 26)
	res := 0

	for _, w := range words {
		flag := true
		copy(temp, charSlice) // 注意这里要用深拷贝，不能用切片切
		for _, c := range []byte(w) {
			idx := c - 'a'
			if temp[idx] > 0 {
				temp[idx]--
			} else {
				flag = false
				break
			}
		}
		if flag == true {
			res += len(w)
		}
	}
	return res
}