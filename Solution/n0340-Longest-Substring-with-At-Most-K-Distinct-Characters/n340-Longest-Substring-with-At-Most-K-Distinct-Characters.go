package main

import ."LeetCode-go/utils"

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	start, res := 0, 0
	mp := map[byte]int{}
	for end := 0; end < len(s); end++ {
		c := s[end]
		mp[c]++

		for len(mp) > k {
			w := s[start]
			mp[w]--
			if mp[w] == 0 {
				delete(mp, w)
			}
			start++
		}
		res = Max(res, end - start + 1)
	}
	return res
}

// 使用数组代替哈希表
func lengthOfLongestSubstringKDistinct1(s string, k int) int {
	start, cnt, res := 0, 0, 0
	mp := make([]int, 256)
	for end := 0; end < len(s); end++ {
		c := s[end] - 'a'
		if mp[c] == 0 {
			cnt++
		}
		mp[c]++

		for cnt > k {
			w := s[start] -  'a'
			mp[w]--
			if mp[w] == 0 {
				cnt--
			}
			start++
		}
		res = Max(res, end - start + 1)
	}
	return res
}
