func stringMatching(words []string) []string {
    var res []string
    for i := range words {
        lps := lps(words[i])
        for j := range words {
            if i == j {
                continue
            }
            if contains(words[j], words[i], lps) {
                res = append(res, words[i])
                break
            }
        }
    }
    return res
}

func lps(s string) []int {
    lps := make([]int, len(s))
    currentIndex := 1
    longestSuffix := 0
    for currentIndex < len(s) {
        if s[currentIndex] == s[longestSuffix] {
            longestSuffix++
            lps[currentIndex] = longestSuffix
            currentIndex++
        } else if longestSuffix == 0 {
            currentIndex++
        } else {
            longestSuffix = lps[longestSuffix - 1]
        }
    }
    return lps
}

func contains(s, t string, lps []int) bool {
    i, j := 0, 0
    for i < len(s) {
        if s[i] == t[j] {
            i++
            j++
            if j == len(t) {
                return true
            }
        } else if j == 0 {
            i++
        } else {
            j = lps[j - 1]
        }
    }
    return false
}