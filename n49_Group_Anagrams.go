package main

import "strings"

//Time: O(n * k), k是最长单词的长度; Space:O(n)
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	temp := make(map[string][]string)

	for _, v := range strs {
		key := getKeyByCount(v)
		temp[key] = append(temp[key], v)
	}

	res := make([][]string, 0, len(temp))
	for _, v := range temp {
		res = append(res, v)
	}
	return res

}

func getKeyByCount(str string) string {
	count := [26]int{}

	for _, v := range str {
		count[v - 'a']++
	}

	var s strings.Builder
	for index, freq := range count {
		if freq != 0 {
			s.WriteString(string(index + 'a'))
			s.WriteString(string(freq))
		}
	}
	return s.String()
}