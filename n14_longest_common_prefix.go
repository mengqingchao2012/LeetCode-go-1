package main

//Time: O(kn)，k是n个字符串最长公共前缀的长度，Space:(1)
func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	}
	first := strs[0]
	for i:= 0; i < len(first); i++ {
		ch := first[i]
		for j := 1; j < size; j++ {
			if i >=len(strs[j]) || ch != strs[j][i] {
				return first[ : i]
			}
		}
	}
	return first
}