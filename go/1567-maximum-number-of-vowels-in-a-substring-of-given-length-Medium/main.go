func maxVowels(s string, k int) int {
    v := 0
    for i := 0; i < k; i++ {
        if vowel(s[i]) {
            v++
        }
    }
    w := v
    for i := k; i < len(s); i++ {
        if vowel(s[i]) {
            w++
        }
        if vowel(s[i - k]) {
            w--
        }
        if w > v {
            v = w
        }
    }
    return v
}

func vowel(b byte) bool {
    return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}