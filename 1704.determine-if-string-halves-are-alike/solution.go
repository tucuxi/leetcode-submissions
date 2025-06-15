func halvesAreAlike(s string) bool {
    vowels := "aeiouAEIOU"
    v := 0
    for _, r := range s[:len(s)/2] {
        if strings.ContainsRune(vowels, r) {
            v++
        }
    }
    for _, r := range s[len(s)/2:] {
        if strings.ContainsRune(vowels, r) {
            v--
        }
    }
    return v == 0
}