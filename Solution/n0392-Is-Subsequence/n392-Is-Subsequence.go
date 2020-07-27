package main

func isSubsequence(s string, t string) bool {
	j, length := 0, len(s)
	if length == 0 { return true }

	for i := 0; i < len(t); i++ {
		if s[j] == t[i] {
			j++
		}
		if j == length {
			return true
		}
	}
	return false
}
