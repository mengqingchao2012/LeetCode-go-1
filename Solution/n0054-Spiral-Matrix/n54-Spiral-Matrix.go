package main

//Time：O(n)；Space：O(1)
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}
	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}

	top, bottom, left, right := 0, row-1, 0, col-1
	res := make([]int, 0, row*col)
	for {
		//遍历上边界
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++ //上边界下移
		if top > bottom {
			break
		}

		//遍历右边界
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right-- //右边界左移
		if left > right {
			break
		}

		//遍历下边界
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom-- //下边界上移
		if top > bottom {
			break
		}

		//遍历左边界
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++ //左边界右移
		if left > right {
			break
		}
	}
	return res
}

//使用闭包
func spiralOrder1(matrix [][]int) []int {
	row := len(matrix)
	if row == 0 {
		return []int{}
	}

	col := len(matrix[0])
	if col == 0 {
		return []int{}
	}

	next := Next(row, col)

	res := make([]int, row*col)
	for idx := range res {
		x, y := next()
		res[idx] = matrix[x][y]
	}
	return res
}

func Next(row, col int) func() (int, int) {
	top, bottom := 0, row-1
	left, right := 0, col-1
	x, y := 0, -1
	dx, dy := 0, 1

	return func() (int, int) {
		x += dx
		y += dy
		switch {
		case y+dy > right:
			top++
			dx, dy = 1, 0
		case x+dx > bottom:
			right--
			dx, dy = 0, -1
		case y+dy < left:
			bottom--
			dx, dy = -1, 0
		case x+dx < top:
			left++
			dx, dy = 0, 1
		}
		return x, y
	}
}
