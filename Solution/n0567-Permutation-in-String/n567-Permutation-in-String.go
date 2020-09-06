package main

func checkInclusion(s1 string, s2 string) bool {
	start, match := 0, 0
	mp := map[byte]int{}

	for idx := range s1 {
		mp[s1[idx]]++
	}
	size := len(mp)
	patternSize := len(s1) - 1

	for end := range s2 {
		c := s2[end]
		if _, ok := mp[c]; ok {
			mp[c]--
			if mp[c] == 0 {
				match++
			}
		}

		if match == size { return true }

		if end - start >= patternSize {
			w := s2[start]
			if _, ok := mp[w]; ok {
				if mp[w] == 0 {
					match--
				}
				mp[w]++
			}
			start++
		}
	}
	return false
}