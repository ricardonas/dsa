func groupAnagrams(strs []string) [][]string {
    
	hasher := make(map[string][]string)

	for _, word := range strs {
		var characters []string = strings.Split(word, "")
		sort.Strings(characters)
		key := strings.Join(characters, "")

		if _, exists := hasher[key]; exists {
			hasher[key] = append(hasher[key], word)
		} else {
			hasher[key] = []string{word}
		}
	}

	result := [][]string{}

	for _, v := range hasher {
		result = append(result, v)
	}

	return result

}