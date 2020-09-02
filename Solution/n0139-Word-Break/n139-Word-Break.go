package main

func wordBreak(s string, wordDict []string) bool {
	dic := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		dic[v] = struct{}{}
	}

	dp := make([]bool, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if _, ok := dic[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
