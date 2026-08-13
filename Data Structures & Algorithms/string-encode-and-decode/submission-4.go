type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _, word := range strs {
		for j, char := range word {
			if j > 0 {
				builder.WriteString("&")
			}
			builder.WriteString(strconv.Itoa(int(char)))
		}
		builder.WriteString("#")
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	words := []string{}

	parts := strings.Split(encoded, "#")
	parts = parts[:len(parts)-1]

	for _, part := range parts {
		if part == "" {
			words = append(words, "")
			continue
		}

		var runeSlice []rune

		charCodes := strings.Split(part, "&")
		for _, code := range charCodes {
			runeValue, _ := strconv.Atoi(code)
			runeSlice = append(runeSlice, rune(runeValue))
		}
		words = append(words, string(runeSlice))

	}
	return words
}
