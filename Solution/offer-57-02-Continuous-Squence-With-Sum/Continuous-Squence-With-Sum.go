package main

func findContinuousSequence(target int) [][]int {
	if target < 3 {
		return nil
	}

	totalNum := make([]int, 0, target)
	for i := 0; i < target; i++ {
		totalNum = append(totalNum, i)
	}

	small, big := 1, 2
	mid := (1 + target) >> 1
	curSum := small + big
	res := [][]int{}

	for small < mid {
		if curSum == target {
			res = append(res, totalNum[small:big + 1])
		}

		for small < mid && curSum > target {
			curSum -= small
			small++

			if curSum == target {
				res = append(res, totalNum[small:big + 1])
			}
		}

		big++
		curSum += big
	}
	return res
}
