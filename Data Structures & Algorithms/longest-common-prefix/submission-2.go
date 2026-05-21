func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
    substr := []rune(strs[0])

	for _, word := range strs[1:] {
		// accout for empty substr
		if len(substr) < 1 {
			return ""
		}
		wordRunes := []rune(word)

		substr = substr[:int(math.Min(float64(len(substr)), float64(len(wordRunes))))]

		for i := range len(substr) {
			if wordRunes[i] != substr[i] {
				substr = substr[:i]
				break
			}
		}
	}

	return string(substr)
}
