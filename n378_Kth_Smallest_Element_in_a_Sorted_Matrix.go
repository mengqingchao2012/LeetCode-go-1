package main

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix) - 1
	left, right := matrix[0][0], matrix[n][n] + 1
	for left < right {
		mid := left + ((right - left) >> 1)
		if nLessMids(&matrix, n, mid) >= k {
			right = mid
		} else  {
			left = mid + 1
		}
	}
	return left
}

func nLessMids(matrix *[][]int, row, mid int) int {
	count := 0
	for i := 0; i <= row; i++ {
		if (*matrix)[i][row] <= mid {
			count += row + 1
		} else {
			for j:=0; j <= row && (*matrix)[i][j] <= mid; j++{
				count++
			}
		}
	}
	return count
}

//func main() {
//	var matrix = [][]int{
//		{1, 2},
//		{1, 3},
//	}
//
//	res := kthSmallest(matrix, 2)
//	fmt.Println(res)
//}
