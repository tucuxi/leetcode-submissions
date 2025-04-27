func vowelStrings(words []string, queries [][]int) []int {
    res := make([]int, 0, len(queries))
    pre := prefixSum(words)
    for _, q := range queries {
        res = append(res, pre[q[1] + 1] - pre[q[0]])
    }
    return res
}

func prefixSum(words []string) []int {
    res := make([]int, 1, len(words) + 1)
    count := 0
    for _, w := range words {
        if len(w) > 0 && isVowel(w[0]) && isVowel(w[len(w) - 1]) {
            count++
        }
        res = append(res, count)
    }
    return res
}

func isVowel(ch byte) bool {
    return strings.IndexByte("aeiou", ch) != -1
}