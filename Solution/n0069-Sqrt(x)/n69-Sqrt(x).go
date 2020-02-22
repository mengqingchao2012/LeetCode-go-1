package main

//二分法
func mySqrt(x int) int {
	//将所有值转换为 int64 是为了防止平方运算时溢出
	num := int64(x)
	var low, high int64 = 0, num

	for low <= high {
		mid := low + ((high - low) >> 1)
		if mid*mid < num {
			low = mid + 1
		} else if mid*mid > num {
			high = mid - 1
		} else {
			return int(mid)
		}
	}
	return int(high) //注意返回值是 high
}

//牛顿迭代
func mySqrt1(x int) int {
	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}
