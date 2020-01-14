package main

import (
	"sort"
)

func frequencySort(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	temp := make([]int, 128)

	for i := range s { //遍历数组，将每个字符出现的次数用数组存储起来
		temp[s[i]]++ //注意：字符串中可能有数字和字符，且区分大小写，不能用 s[i] - 'a'来算下标
	}

	ss := make([]string, 128) //注意容量取到128（ASCII码到128）, 同时有数字，大小写是个坑
	for idx, val := range temp {
		if val == 0 {
			continue
		}
		tp := make([]byte, 0, 128)
		//将每个出现的字符展开成string，保存在ss中
		for val > 0 {
			tp = append(tp, byte(idx))
			val--
		}
		ss = append(ss, string(tp))
	}

	sort.Slice(ss, func(i, j int) bool {
		if len(ss[i]) == len(ss[j]) { //如果字符长度相等，则比较其ASCII码
			return ss[i] < ss[j]
		}
		return len(ss[i]) > len(ss[j]) //否则比较长度，即频率
	})

	var res string
	for _, v := range ss {
		if v != "" {
			res += v
		}
	}
	return res
}

//func main() {
//	fmt.Println(frequencySort("Aabb"))
//}
