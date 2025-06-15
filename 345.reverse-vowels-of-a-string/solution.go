func reverseVowels(s string) string {
    vowels := map[byte]bool{
        'A': true,
        'a': true,
        'E': true,
        'e': true,
        'I': true,
        'i': true,
        'O': true,
        'o': true,
        'U': true,
        'u': true,
    }
    r := make([]byte, len(s))
    for i, j := 0, len(r) - 1;; i, j = i + 1, j - 1 {
        for i <= j && !vowels[s[i]] {
            r[i] = s[i]
            i++
        }
        for i <= j && !vowels[s[j]] {
            r[j] = s[j]
            j--
        }
        if i > j {
            break
        }
        r[i], r[j] = s[j], s[i]
    }
    return string(r)
}

func containsRune(s string, r rune) bool {
    for _, c := range s {
        if c == r {
            return true
        }
    }
    return false
}