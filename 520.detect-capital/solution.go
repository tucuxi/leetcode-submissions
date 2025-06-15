func detectCapitalUse(word string) bool {
    uc, lc := 0, 0
    for i := range word {
        if word[i] >= 'A' && word[i] <= 'Z' {
            uc++
        } else {
            lc++
        }
    }
    return uc == len(word) || lc == len(word) || uc == 1 && word[0] >= 'A' && word[0] <= 'Z'
}