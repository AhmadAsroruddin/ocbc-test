package main

func hotDog(wordDict []string, firstWord string, lastWord string) []string {
	var result []string

	if wordDict[0] != firstWord {
		wordDict = append([]string{firstWord}, wordDict...)
	}

	for _, word := range wordDict {
		result = append(result, word)

		if word == lastWord {
			result = append(result, word)
			return result
		} else if wordCheck(lastWord, word) {
			result = append(result, lastWord)
			return result
		}
	}
	return []string{"<no way>"}
}

func wordCheck(input string, target string) bool {
	setInput := make(map[rune]bool)
	setTarget := make(map[rune]bool)

	for _, word := range input {
		setInput[word] = true
	}

	for _, word := range target {
		setTarget[word] = true
	}

	counter := 0
	for word := range setInput {
		if !setTarget[word] {
			counter++
		}
	}
	if counter > 1 {
		return false
	}
	return true
}
