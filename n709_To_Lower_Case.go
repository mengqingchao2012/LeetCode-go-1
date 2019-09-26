package main

func toLowerCase(str string) string {
	length := len(str)
	if length == 0 {
		return str
	}

	words := []byte(str)

	for i := 0; i < length; i++ {
		if words[i] >= 65 && words[i] <= 90 {
			words[i] += 32
		}
	}
	return string(words)
}
