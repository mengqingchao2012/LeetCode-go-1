package main

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for index, i := range words {
		start, end := 0, len(i)-1
		w := []byte(i)
		for start <= end {
			w[start], w[end] = w[end], w[start]
			start++
			end--
		}
		words[index] = string(w)
	}
	result := strings.Join(words, " ")
	return result
}
