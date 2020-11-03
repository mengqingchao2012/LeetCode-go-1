package main

//Time: O(kn)，k是n个字符串最长公共前缀的长度，Space:(1)
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	first := strs[0]
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || first[i] != strs[j][i] {
				return first[:i]
			}
		}
	}
	return first
}
