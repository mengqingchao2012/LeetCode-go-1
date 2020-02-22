package main

import "strings"

func reverseWordsI(s string) string {
	if s == "" {
		return s
	}
	temp := strings.Split(s, " ")
	var res []string
	for _, v := range temp {
		if v != "" {
			res = append(res, v)
		}
	}
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return strings.Join(res, " ")
}

//func main() {
//	reverseWordsI("a good   example")
//}
