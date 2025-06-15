func vowelStrings(words []string, left int, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        w := words[i]
        if isVowel(w[0]) && isVowel(w[len(w) - 1]) {
            count++
        }
    }
    return count
}

func isVowel(b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}