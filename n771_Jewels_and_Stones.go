package main

func numJewelsInStones(J string, S string) int {
	isJewe := make(map[rune]bool, len(J)) //注意用的是rune，不是byte
	for _, i := range J {
		isJewe[i] = true
	}

	count := 0
	for _, j := range S {
		if isJewe[j] {
			count++
		}
	}

	return count
}