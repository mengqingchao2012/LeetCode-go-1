package main

//方法一：Time：O(1)，Space：O(1)
func isHappy(n int) bool {
	if n < 0 {
		return false
	}

	//借助一个辅助map，将每次计算的结果加入map中，如果出现重复，就返回false
	mp := make(map[int]bool)
	mp[n] = true

	for n != 1 {
		n = transfer(n)
		if mp[n] == true {
			return false
		} else {
			mp[n] = true
		}
	}
	return true
}

//方法二：Time：O(1)，Space：O(1)
func isHappy1(n int) bool {
	if n < 0 {
		return false
	}

	//借助快慢指针的思想，快指针每一轮迭代两次，慢指针迭代一次
	fast, slow := n, n
	for {
		fast = transfer(transfer(fast))
		slow = transfer(slow)
		if fast == 1 {return true} //如果快指针等于1，则可以直接返回
		if fast == slow {return false} //如果慢指针等于快指针，则返回false
	}
	return false
}

func transfer(n int) int {
	sum := 0
	for n != 0 {
		a := n % 10
		sum += a * a
		n /= 10
	}
	return sum
}
