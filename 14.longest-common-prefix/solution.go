func longestCommonPrefix(strs []string) string {
    for i := range strs[0] {
        for j := 1; j < len(strs); j++ {
            if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}