package main

// 方法一：朴素筛
func countPrimes(n int) int {
	isComposite := make([]bool, n + 1)
	cnt := 0

	for i := 2; i < n; i++ {
		if !isComposite[i] {
			cnt++
		}
		for j := i + i; j < n; j = j + i {
			isComposite[j] = true
		}
	}
	return cnt
}

// 方法二：埃氏筛
func countPrimes1(n int) int {
	isComposite := make([]bool, n + 1)
	cnt := 0

	for i := 2; i < n; i++ {
		if !isComposite[i] {
			cnt++
			for j := i + i; j < n; j = j + i {
				isComposite[j] = true
			}
		}
	}
	return cnt
}