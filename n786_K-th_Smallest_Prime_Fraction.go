package main

import "fmt"

func kthSmallestPrimeFraction(A []int, K int) []int {
	var left float64= 0.0
	var right float64= 1.0
	size := len(A)

	for left < right {
		var mid = float64(left+right) / 2
		var max float64 = 0
		p, q, j, count := 0, 0, 1, 0

		for i := 0; i < size - 1; i++ {
			for j < size && float64(A[i]) > float64(A[j]) * mid {
				j++
			}
			if size == j {
				break
			}
			count += size - j

			temp := float64(A[i]) / float64(A[j])
			if temp > max {
				max = temp
				p = i
				q = j
			}
		}

		if count == K {
			return []int{A[p], A[q]}
		} else if count > K {
			right = mid
		} else {
			left = mid
		}
	}
	return []int{}
}

func main() {
	A := []int{1, 2, 3, 5}
	res := kthSmallestPrimeFraction(A, 3)
	fmt.Println(res)
}
