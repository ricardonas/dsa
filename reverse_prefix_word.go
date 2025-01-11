func reversePrefix(word string, ch byte) string {
	characters, charIndex := []byte(word), -1

	for index, value := range characters {
		if value == ch {
			charIndex = index
			break
		}
	}

	if charIndex == -1 {
		return word
	}

	p1, p2 := 0, charIndex

	for p1 < p2 {
		characters[p1], characters[p2] = characters[p2], characters[p1]
		p1, p2 = p1+1, p2-1
	}

	return string(characters)

}

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

	reversedWord = []byte(word)

	for p1 := 0; p1 < p2; p1++ {
		reversedWord[p1], reversedWord[p2] = reversedWord[p2], reversedWord[p1]
		p2 -= 1
	}

	return reversedWord
}