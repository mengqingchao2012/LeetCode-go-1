package main

func findAnagrams(s string, p string) []int {
	start, matched := 0, 0
	mp := map[byte]int{}
	res := []int{}

	for i := range p {
		mp[p[i]]++
	}
	size := len(mp)
	patternSize := len(p) - 1

	for end := range s {
		w := s[end]
		if _, ok := mp[w]; ok {
			mp[w]--
			if mp[w] == 0 {
				matched++
			}
		}

		if matched == size {
			res = append(res, start)
		}

		if end - start >= patternSize {
			c := s[start]
			if _, ok := mp[c]; ok {
				if mp[c] == 0 {
					matched--
				}
				mp[c]++
			}
			start++
		}
	}
	return res
}