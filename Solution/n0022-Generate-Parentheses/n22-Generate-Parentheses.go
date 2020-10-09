package main

// 递归
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}

	left, right := n, n
	res := []string{}
	recursive(&res, "", left, right)
	return res
}

func recursive(res *[]string, cur string, left, right int) {
	// 递归退出条件：括号全部用完，得到一组合法解
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	// 先写左括号，再写右括号
	if left > 0 {
		recursive(res, cur + "(", left - 1, right)
	}
	if right > left {
		recursive(res, cur + ")", left, right - 1)
	}
}