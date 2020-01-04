package main

//方法一：使用位运算，参考136题
//难点：使用异或最终只能得到一个数，然而现在需要找两个数，如果能把这两个数分到两堆里分别求异或，就可以得到这两个数了
//分离两个数：两个数不相同，其二进制表示必然在某一位上不相同（一个0一个1），遍历数组并将所有结果异或起来，然后
//找到最终结果里第一个出现1的位，就可以通过这一位是0还是1来将数字分成两组
func singleNumber(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	/*
	* 等价于
	* mask := 1
	* for (xor & mask) == 0 {mask <<= 1}
	*/
	mask := xor & (-xor) //找到第一个不是0的位，注意是 & 运算

	x, y := 0, 0
	for _, v := range nums {
		//一边遍历，一边进行位运算
		if (v & mask) == 0 { //注意是 & 运算
			x ^= v
		} else {
			y ^= v
		}
	}
	return []int{x, y}
}

//方法二
func singleNumber1(nums []int) []int {
	if len(nums) == 0 {
		return []int{0, 0}
	}
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}

	res := make([]int, 0, 2)
	for k, v := range mp {
		if v == 1{
			res = append(res, k)
		}
	}
	return res
}
