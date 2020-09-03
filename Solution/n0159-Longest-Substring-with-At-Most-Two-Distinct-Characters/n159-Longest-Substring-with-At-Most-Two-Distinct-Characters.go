package main

import ."LeetCode-go/utils"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	start, cnt, res := 0, 0, 0
	mp := make([]int, 256)

	for end := 0; end < len(s); end++ {
		c := s[end] - 'a'
		if mp[c] == 0 { cnt++ }
		mp[c]++

		for cnt > 2 {
			w := s[start] - 'a'
			mp[w]--
			if mp[w] == 0 { cnt-- }
			start++
		}

		res = Max(res, end - start + 1)
	}
	return res
}