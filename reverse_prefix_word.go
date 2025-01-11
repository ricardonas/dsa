// Two pointers.
func reversePrefix(word string, ch byte) string {
	p2 := -1

	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			p2 = i
			break
		}
	}

	if p2 == -1 {
		return word
	}

	reversedWord := []byte(word)

	for p1 := 0; p1 < p2; p1++ {
		reversedWord[p1], reversedWord[p2] = reversedWord[p2], reversedWord[p1]
		p2 -= 1
	}

	return string(reversedWord)
}