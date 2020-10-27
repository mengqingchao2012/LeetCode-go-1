package main

import ."LeetCode-go/utils"

// 方法一：动态规划
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true
	}

	start, maxLen := 0, 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if i + 1 == j {
				isPalindrome[i][j] = s[i] == s[j]
			} else {
				isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i + 1][j - 1]
			}

			if isPalindrome[i][j] && j - i + 1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start:start + maxLen]
}

// 方法二：中心扩展
func longestPalindrome1(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	start, maxLen := 0, 0
	for i := 0; i < n; i++ {
		l1 := expand(s, i, i)
		l2 := expand(s, i, i + 1)
		tmpLen := Max(l1, l2)

		if tmpLen > maxLen {
			maxLen = tmpLen
			start = i - ((tmpLen - 1) >> 1)
		}
	}
	return s[start:start + maxLen]
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// (right - 1) - (left + 1) + 1
	return right - left - 1
}