func maxFreqSum(s string) int {
    var f [26]int
    var maxConsonant, maxVowel int

    for _, r := range s {
        i := r - 'a'
        f[i]++
        if isVowel(r) {
            maxVowel = max(maxVowel, f[i])
        } else {
            maxConsonant = max(maxConsonant, f[i])
        }
    }
    return maxConsonant + maxVowel
}

func isVowel(r rune) bool {
    return strings.IndexRune("aeiou", r) != -1
}