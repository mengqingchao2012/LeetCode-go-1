package main

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex + 1)
	res[0] = 1
	for i := 0; i <= rowIndex; i++ {
		res[i] = 1
		for j := i - 1; j >= 1; j-- {
			res[j] = res[j - 1] + res[j]
		}
	}
	return res
}