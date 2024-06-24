package main

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)

	for _, word := range wordDict {
		wordMap[word] = true
	}
	return wordBreakHelper(s, wordMap)
}

func wordBreakHelper(s string, wordDict map[string]bool) []string {
	var result []string

	if wordDict[s] {
		result = append(result, s)
	}

	for i := 1; i < len(s); i++ {
		start := s[:i]
		if wordDict[start] {
			suffix := s[i:]
			suffixWays := wordBreakHelper(suffix, wordDict)
			for _, way := range suffixWays {
				result = append(result, start+" "+way)
			}
		}
	}
	return result
}
