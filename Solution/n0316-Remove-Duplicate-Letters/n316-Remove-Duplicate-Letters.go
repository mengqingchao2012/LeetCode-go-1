package main

func removeDuplicateLetters(s string) string {
	res := make([]rune, 0, len(s)) // 最终的结果集
	inRes := map[rune]bool{} // 表示当前字符是否已经存在结果集中
	lastFind := map[rune]int{} // 表示当前字符在原字符串中最后出现的位置

	for i, v := range s {
		lastFind[v] = i
	}

	for i, v := range s {
		if exist := inRes[v]; exist { // 如果当前字符已经存在在最终结果集里，继续循环
			continue
		}
		// 对于结果集中的最后一个字符，如果该字符的字典序大于当前字符，且该字符在后面还出现过，则去掉该字符，加入当前字符
		for len(res) != 0 && res[len(res) - 1] > v && lastFind[res[len(res) - 1]] > i {
			inRes[res[len(res) - 1]] = false
			res = res[:len(res) - 1]
		}
		res = append(res, v)
		inRes[v] = true
	}
	return string(res)
}
