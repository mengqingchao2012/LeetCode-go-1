package main

// 方法一
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := make([]int, 0, len(T))

	for i := 0; i < len(T); i++ {
		for len(stack) != 0 && T[stack[len(stack) - 1]] < T[i] {
			idx := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res[idx] = i - idx
		}
		stack = append(stack, i)
	}

	for _, v := range stack {
		res[v] = 0
	}
	return res
}

// 方法二
func dailyTemperatures1(T []int) []int {
	res := make([]int, len(T))
	res[len(res) - 1] = 0
	for i := len(T) - 2; i >= 0; i-- {
		step := i + 1
		for T[step] <= T[i] && res[step] != 0 {
			step += res[step]
		}
		if T[step] > T[i] {
			res[i] = step - i
		} else {
			res[i] = 0
		}
	}
	return res
}
