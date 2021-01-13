package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}

	i := 0
	res := []string{}
	for i < n {
		j := i + 1
		for j < n && nums[j] - nums[j - 1] == 1 {
			j++
		}
		if i == j - 1 {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(nums[i]), strconv.Itoa(nums[j - 1])))
		}
		i = j
	}

	return res
}
