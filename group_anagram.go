func groupAnagrams(strs []string) [][]string {
    
	hasher := make(map[string][]string)

	for _, word := range strs {
		var characters []string = strings.Split(word, "")
		sort.Strings(characters)
		key := strings.Join(characters, "")
		hasher[key] = append(hasher[key], word)
	}

	//result := [][]string{}

	result := make([][]string, 0, len(hasher))
	
	for _, v := range hasher {
		result = append(result, v)
	}

	return result

}