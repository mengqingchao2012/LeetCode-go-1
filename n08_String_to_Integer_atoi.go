package main

import (
	"math"
)

func myAtoi(str string) int {
	start, p, length := 0, 0, len(str)
	var negative = false //标记符号位
	words := []byte(str)

	for p < length && words[p] == ' ' { //去掉先导空格
		p++
	}
	if p == length { //如果全是空格，直接返回
		return 0
	}

	if words[p] == '+' { //去掉符号位
		p++
	} else if words[p] == '-' {
		negative = true
		p++
	}

	for p < length && words[p] == '0' { //去掉先导0
		p++
	}
	start = p

	for p < length && words[p] >= '0' && words[p] <= '9' { //遍历计算数字的长度
		p++
	}
	if p == start { //如果没有检测到任何数字，直接返回
		return 0
	}

	if p - start > 10 { //如果数字的长度超过10，则肯定超过了int的取值范围，直接根据符号返回极值
		if negative {
			return math.MinInt32
		} else {
			return math.MaxInt32
		}
	}

	var num int64 = 0 //用long类型来存储，防止超限
	for i := start; i < p; i++ {
		num = num * 10 + int64(words[i] - '0') // 注意一定要写 -"0"
	}
	if negative {
		num = -num
	}

	switch { //根据符号位返回最后的结果
	case num < math.MinInt32:
		return math.MinInt32
	case num > math.MaxInt32:
		return math.MaxInt32
	default:
		return int(num)
	}
}

//func main() {
//	res := myAtoi("2147483648")
//	fmt.Println(res)
//}