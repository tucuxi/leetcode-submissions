func countPrefixSuffixPairs(words []string) int {
    res := 0

    for i, str1 := range words {
        for _, str2 := range words[i + 1 :] {
            if len(str1) <= len(str2) && str1 == str2[: len(str1)] && str1 == str2[len(str2) - len(str1) :] {
                res++
            }
        }
    }
    return res
}