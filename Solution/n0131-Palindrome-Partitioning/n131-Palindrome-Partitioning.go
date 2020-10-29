package main

func partition(s string) [][]string {
	n := len(s)
	if n == 0 {
		return [][]string{}
	}

	res := [][]string{} // 最终的结果集
	elem := make([]string, 0, n) // 存储中间生成的结果
	isPalindrome := make([][]bool, n) // isPalindrome[i][j] 表示子串[i:j]是否是回文子串
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}

	// 遍历字符串，标记出其中所有的子串，见第 5 题，时间复杂度：O(n^2）
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i + 1 == j {
				isPalindrome[i][j] = s[i] == s[j]
			} else {
				isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i + 1][j - 1]
			}
		}
	}

	cutSubString(s, 0, isPalindrome, &res, elem)
	return res
}

func cutSubString(s string, start int, isPalindrome [][]bool, res *[][]string, elem []string) {
	if start >= len(s) { // 如果 start 大于等于字符串长度，说明找到了一组合法的解，加入结果集中
		tmp := make([]string, len(elem))
		copy(tmp, elem)
		*res = append(*res, tmp)
		return
	}

	// 给定任意 start，遍历 [start:]，枚举所有可能的子串
	for end := start; end < len(s); end++ {
		if isPalindrome[start][end] {
			// [start:end] 是回文子串，则加入临时结果集 elem 中
			// 更新 start = end + 1，继续枚举 [end + 1:]
			elem = append(elem, s[start:end + 1])
			cutSubString(s, end + 1, isPalindrome, res, elem)
			// 退递归后要将最后一个加入的元素删掉，继续循环枚举 [start:end + 1]
			elem = elem[:len(elem) - 1]
		}
	}
}
