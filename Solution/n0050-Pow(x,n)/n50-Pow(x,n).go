package main

func myPow(x float64, n int) float64 {
	m := abs(int64(n)) //判断 n 的符号，注意类型限制，n 为 int，如果n是最小的 int 的话，取绝对值后会溢出
	var result float64 = 1

	for m != 0 {
		if (m & 1) == 1 { //如果 m 的最低位值为1
			result *= x //则更新结果
		}
		x *= x  //更新底数
		m >>= 1 //m右移更新最低位
	}

	if n >= 0 {
		return result
	} else {
		return 1 / result
	}
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
