func strStr(haystack string, needle string) int {
    p := prefix(needle)
    i := 0
    j := 0
    for i < len(haystack) {
        for j >= 0 && haystack[i] != needle[j] {
            j = p[j]
        }
        i++
        j++
        if j == len(needle) {
            return i - len(needle)
        }
    }
    return -1       
}

func prefix(w string) []int {
    p := make([]int, len(w) + 1)
    i := 0
    j := -1
    p[i] = j
    for i < len(w) {
        for j >= 0 && w[j] != w[i] {
            j = p[j]
        }
        i++
        j++
        p[i] = j
    }
    return p
}