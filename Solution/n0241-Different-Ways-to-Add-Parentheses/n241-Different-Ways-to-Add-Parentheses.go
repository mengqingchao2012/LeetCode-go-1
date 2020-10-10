package main

import "strconv"

func diffWaysToCompute(input string) []int {
	if n, err := strconv.Atoi(input); err == nil {
		return []int{n}
	}

	res := []int{}

	for idx := range input {
		if input[idx] == '+' || input[idx] == '-' || input[idx] == '*' {
			left := diffWaysToCompute(input[:idx])
			right := diffWaysToCompute(input[idx + 1:])

			for _, l := range left {
				for _, r := range right {
					tmp := 0
					switch input[idx] {
					case '+':
						tmp = l + r
					case '-':
						tmp = l - r
					case '*':
						tmp = l * r
					}
					res = append(res, tmp)
				}
			}
		}
	}
	return res
}