func reverseString(s []byte) {
	for i := range len(s)/2 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
