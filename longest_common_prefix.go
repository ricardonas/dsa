func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	
	sort.Strings(strs)

	index, str1, str2 := 0, strs[0], strs[len(strs)-1]

	for index < len(str1) {
		if str1[index] != str2[index] {
			break;
		}
		index++
	}
	
	if index == 0 {
		return ""
	}

	return str1[0:index]
}
