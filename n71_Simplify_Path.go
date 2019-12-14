package main

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	temp := strings.Split(path, "/")
	for _, v := range temp {
		switch v {
		case ".", "":
		case "..":
			if len(stack) != 0 {
				stack = stack[:len(stack) - 1]
			}
		default :
			stack = append(stack, v)
		}
	}
	return "/" + strings.Join(stack, "/")
}
