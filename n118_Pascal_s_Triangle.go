package main

//杨辉三角形的特点：
//1.每一行的首尾两个元素都是1；
//2.每一行的元素个数是i + 1；
//3.每一行中除了首末元素外，剩下的元素都满足a[i][j] = a[i - 1][j - 1] + a[i - 1][j]
//Time：O(n^2)，Space：O(1)
func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows < 1 {
		return res
	}

	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, temp)
	}
	return res
}
