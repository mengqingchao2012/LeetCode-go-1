package main

import ."LeetCode-go/utils"

func characterReplacement(s string, k int) int {
	start, maxCount, res := 0, 0, 0
	mp := make([]int, 128)

	for end := 0; end < len(s); end++ {
		c := s[end] - 'A'
		mp[c]++
		maxCount = Max(maxCount, mp[c])

		for end - start + 1 - maxCount > k {
			w := s[start] - 'A'
			mp[w]--
			start++
		}
		res = Max(res, end - start + 1)
	}
	return res
}
