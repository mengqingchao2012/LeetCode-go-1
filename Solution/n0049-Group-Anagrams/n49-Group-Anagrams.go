package main

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	tmp := map[string][]string{}

	for _, s := range strs {
		key := getKeyByCount(s)
		tmp[key] = append(tmp[key], s)
	}

	for _, v := range tmp {
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
	for w, freq := range count {
		s.WriteString(string(w + 'a'))
		s.WriteString(string(freq))
	}
	return s.String()
}

func groupAnagrams1(strs []string) [][]string {
	res := [][]string{}
	tmp := map[string][]string{}

	for _, s := range strs {
		newStr := []byte(s)
		sort.Slice(newStr, func(i, j int) bool {
			return newStr[i] < newStr[j]
		})
		tmp[string(newStr)] = append(tmp[string(newStr)], s)
	}

	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}
