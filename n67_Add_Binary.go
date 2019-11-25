package main

import (
	"LeetCode-go/utils"
	"strconv"
)

func addBinary(a string, b string) string {
	la := len(a)
	lb := len(b)
	temp := make([]int, 0, utils.Max(la, lb)+1)
	var res string

	i, j := la-1, lb-1
	var carry = 0
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}
		temp = append(temp, carry%2)
		carry /= 2
	}

	for i := len(temp) - 1; i >= 0; i-- {
		res += strconv.Itoa(temp[i])
	}

	return res
}
