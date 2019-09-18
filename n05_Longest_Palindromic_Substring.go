package main

//解法一：Time:O(n^2)，Space:O(n^2)
func longestPalindrome1(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	start, max := 0, 0
	res := make([][]bool, length, length)
	for i := 0; i < length; i++ {
		res[i] = make([]bool, length)
	}

	for i := length - 1; i >= 0; i-- {
		for j := i; j < length; j++ {
			if i == j {
				res[i][j] = true
			} else if i + 1 == j {
				res[i][j] = s[i] == s[j]
			} else {
				res[i][j] = s[i] == s[j] && res[i + 1][j - 1]
			}

			if res[i][j] && j - i + 1 > max {
				start = i
				max = j - i + 1
			}
		}
	}

	return s[start : start + max]
}

//解法二：中心扩展法：Time：O(n^2)，Space：O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expand(i, i, s)
		len2 := expand(i, i + 1, s)
		len := max(len1, len2)

		if len > maxLen {
			start = i - ((len - 1) >> 1)
			maxLen = len
		}
	}
	return s[start : start + maxLen]
}

func expand(left, right int, s string) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}