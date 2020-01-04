package main

//Time：O(n)，Space：O(k)
//遍历两次字符串，第一次统计字符出现的次数，第二次查找第一个统计次数为1的字符下标
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}

	count := make([]int, 26) //用来统计字符出现的次数
	for i := 0; i < len(s); i++ {
		count[s[i] - 'a']++
	}
	for i := 0; i < len(s); i++ {
		if count[s[i] - 'a'] == 1 {
			return i
		}
	}
	return -1
}