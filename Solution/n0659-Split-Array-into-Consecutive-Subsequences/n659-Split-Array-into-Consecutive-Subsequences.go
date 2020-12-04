package main

func isPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	countMap := map[int]int{}
	resMap := map[int]int{}

	for _, v := range nums {
		countMap[v]++
	}

	for _, v := range nums {
		if countMap[v] == 0 {
			continue
		}

		if _, ok := resMap[v - 1]; ok {
			resMap[v - 1]--
			resMap[v]++
			countMap[v]--
		} else {
			if resMap[v + 1] == 0 || resMap[v + 2] == 0 {
				return false
			} else {
				resMap[v]--
				resMap[v + 1]--
				resMap[v + 2]--
				countMap[v + 2]++
			}
		}
	}
	return true
}