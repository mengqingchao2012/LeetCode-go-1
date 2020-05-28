package main

func StrStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	if l1 == 0 && l2 == 0 {
		return 0
	}
	if l1 == 0 {
		return -1
	}
	if l2 == 0 {
		return 0
	}

	for i := 0; i <= l1-l2; i++ { // 注意这里，l1-l2要取等号，不然在 haystack 和 needle 长度相同时会出错
		j, k := 0, i
		for j < l2 && k < l1 && haystack[k] == needle[j] {
			j++
			k++
		}
		if j == l2 {
			return i
		}
	}
	return -1
}
