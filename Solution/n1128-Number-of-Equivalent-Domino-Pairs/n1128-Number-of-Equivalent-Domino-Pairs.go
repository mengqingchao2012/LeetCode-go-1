package main

func numEquivDominoPairs(dominoes [][]int) int {
	n := len(dominoes)
	if n <= 1 {
		return 0
	}

	cnt := [100]int{}
	ans := 0
	for i := 0; i < n; i++ {
		a, b := dominoes[i][0], dominoes[i][1]
		if a > b {
			a, b = b, a
		}
		tmp := a * 10 + b
		ans += cnt[tmp]
		cnt[tmp]++
	}
	return ans
}
