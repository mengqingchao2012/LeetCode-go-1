package main

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	start, matched, subStart := 0, 0, 0
	min_size, patternSize := len(s) + 1, len(t)
	mp := map[byte]int{}

	for idx := range t {
		mp[t[idx]]++
	}

	for end := range s {
		c := s[end]
		if _, ok := mp[c]; ok {
			mp[c]--
			// 注意，因为可能出现两个连续重复字符的情况，如 aa，而只需要匹配一个 a 即可，
			// 因此这里只要匹配过一次该字符，matched 就递增，同时下面触发收缩窗口的条件也相应的要变为
			// matched == len(pattern)，
			if mp[c] >= 0 {
				matched++
			}
		}

		for matched == patternSize {
			if min_size > end - start + 1 {
				min_size = end - start + 1
				subStart = start
			}

			w := s[start]
			if _, ok := mp[w]; ok {
				if mp[w] == 0 {
					matched--
				}
				mp[w]++
			}
			start++
		}
	}
	if min_size > len(s) {
		return ""
	}
	return s[subStart : subStart + min_size]
}
