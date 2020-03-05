package main

import ."LeetCode-go/utils"

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)
	i := 0
	for candies != 0 {
		res[i % num_people] += Min(i + 1, candies)
		candies -= Min(i + 1, candies)
		i++
	}
	return res
}
