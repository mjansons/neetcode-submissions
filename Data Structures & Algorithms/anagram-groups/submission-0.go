import "slices"

func groupAnagrams(strs []string) (final [][]string) {
    if len(strs) == 1 {
        return [][]string{strs}
    }
    groups := make(map[string][]string)

    for _, v := range strs {
        sortedWord := []rune(v)
		slices.Sort(sortedWord)
		sortedWordString := string(sortedWord)

		value, exists := groups[sortedWordString]
		if exists {
			groups[sortedWordString] = append(value, v)
			continue
		}

		groups[sortedWordString] = []string{v}

    }

    for _, v := range groups {
		final = append(final, v)
    }

    return final

}
