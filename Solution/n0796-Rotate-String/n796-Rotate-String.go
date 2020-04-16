package main

func rotateString(A string, B string) bool {
	l1, l2 := len(A), len(B)
	if l1 != l2 {
		return false
	}

	AA := A + A
	n, nn := l1, 2*l1
	for start := 0; start <= nn; start++ { // 注意等号不能漏
		i, j := start, 0
		for i < nn && j < n && AA[i] == B[j] {
			i++
			j++
		}
		if j == n {
			return true
		}
	}
	return false
}