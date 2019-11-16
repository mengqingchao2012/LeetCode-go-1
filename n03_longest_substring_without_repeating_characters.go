package main

import "LeetCode-go/utils"

//方法一：哈希表：Time：O(n)，Space:O(n)
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	mp := make(map[byte]int)
	ans := 0
	for start, end := 0, 0; end < size; end++ {
		if _, err := mp[s[end]]; err != false {
			start = utils.Max(mp[s[end]], start)
		}
		ans = utils.Max(end - start + 1, ans)
		mp[s[end]] = end + 1
	}
	return ans
}

//方法二：数组：Time：O(n)，Space:O(n)
func lengthOfLongestSubstring1(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	store := [128]int{} //与方法一思路相同，使用数组来代替哈希表进行优化
	ans := 0
	for start, end := 0, 0; end < size; end++ {
		if store[s[end]] != 0 {
			start = utils.Max(store[s[end]], start)
		}
		ans = utils.Max(end - start + 1, ans)
		store[s[end]] = end + 1
	}
	return ans
}

