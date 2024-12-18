// This is not my first attempt at solving the problem.
func longestCommonPrefix(strs []string) string {
    i := 0

    str1 := strs[0]
    str2 := strs[len(strs) - 1]

    for i < len(str1) {
        if str1[i] == str2[i] {
            i++
        } else {
            break;
        }
    }

    if i == 0 {
        return ""
    }
    
    return str1[0:i]
}
