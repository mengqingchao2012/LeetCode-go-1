package main

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x > 0 { //边界条件判断：负数以及能被10整除的数都不是回文
		return false
	}

	temp := x
	var res int64 = 0 //注意res应定义为int64，反转数字后可能会超过int的范围
	for temp > 0 {
		num := temp % 10
		res = res*10 + int64(num)
		temp /= 10
	}

	return res == int64(x)
}
