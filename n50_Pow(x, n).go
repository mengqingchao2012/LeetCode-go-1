package main

//要点：注意边界条件的划分
//1.任意数的 0 次幂都为1
//2.x的-n次幂等于x的n次幂的倒数，此时需注意，x若为0， 则会出现异常结果，此处需要声明

//Time：O(logn), Space:(1)
func myPow(x float64, n int) float64 {
	var res float64
	switch {
	case n == 0:
		res = 1
	case n < 0 && x == 0:
		res = 0
	default:
		if n > 0 {
			res = countPow(x, n)
		} else {
			res = 1 / countPow(x, -n)
		}
	}
	return res
}

//二分法：将n写成二进制的形式，则x的n次幂可以进行展开
//如：x^11 = x^8 * x^2 * x^1, 即：11写成二进制是1011
func countPow(x float64, n int) float64 {
	res := 1.0
	for n != 0 {
		if (n & 1) == 1 { //如果当前进制位的值为1，则结果乘以x
			res *= x
		}
		x *= x //x一直保持翻倍递增
		n >>= 1
	}
	return res
}
