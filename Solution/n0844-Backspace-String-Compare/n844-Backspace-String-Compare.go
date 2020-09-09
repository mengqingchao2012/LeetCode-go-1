package main

func backspaceCompare(S string, T string) bool {
	idx1, idx2 := len(S) - 1, len(T) - 1
	for idx1 >= 0 || idx2 >= 0 {
		i := getNextValidIndex(S, idx1)
		j := getNextValidIndex(T, idx2)

		switch {
		case i < 0 && j < 0:
			return true
		case i < 0 || j < 0:
			return false
		case S[i] != T[j]:
			return false
		}

		idx1 = i - 1
		idx2 = j - 1
	}
	return true
}

func getNextValidIndex(s string, idx int) int {
	count := 0
	// 注意：带有 break 语句的条件分支块最好别用 switch 语句，跳转可能会出错
	for idx >= 0 {
		if s[idx] == '#'{
			count += 1
		} else if count > 0 {
			count -= 1
		} else {
			break
		}
		idx--
	}
	return idx
}
