package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compressString(S string) string {
	if len(S) == 0 { // 注意判断S为空的情况
		return S
	}
	count := 0
	word := S[0]
	res := []string{}
	for _, v := range []byte(S) {
		if v != word { // 如果当前字符不等于上一个字符，则将上一个字符和其出现次数加入辅助数组中存储，并更新count和word
			res = append(res, string(word), strconv.Itoa(count))
			count = 1
			word = v
		} else {
			count++
		}
	}
	res = append(res, string(word), strconv.Itoa(count)) // 注意最后退出循环后，还要将最后一组count和word加入结果集中
	if len(res) < len(S) {
		return strings.Join(res, "")
	} else {
		return S
	}
}

func main() {
	fmt.Println(compressString("aabcccccaa"))
}
