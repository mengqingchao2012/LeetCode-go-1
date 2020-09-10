package main

// 方法一：快慢指针
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = next(slow)
		fast = next(next(fast))

		// 注意：要先判断 fast == 1，因为如果 n = 1 的话，先判断 slow == fast 会直接退出 false
		if fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
	return false
}

func next(n int) int {
	sum := 0
	for n != 0 {
		tmp := n % 10
		sum += tmp * tmp
		n /= 10
	}
	return sum
}

// 方法二：辅助哈希表
func isHappy1(n int) bool {
	//借助一个辅助map，将每次计算的结果加入map中，如果出现重复，就返回false
	mp := make(map[int]bool)
	mp[n] = true

	for n != 1 {
		n = next(n)
		if mp[n] == true {
			return false
		} else {
			mp[n] = true
		}
	}
	return true
}