package main

func flipAndInvertImage(A [][]int) [][]int {
	n := len(A)
	if n == 0 {
		return A
	}

	c := len(A[0])
	for i := 0; i < n; i++ {
		a := A[i]
		for j := 0; j < (c + 1)/2; j++ { // 边界 j < (c + 1) / 2
			a[j], a[c - j - 1] = a[c - j - 1] ^ 1, a[j] ^ 1
		}
	}
	return A
}
