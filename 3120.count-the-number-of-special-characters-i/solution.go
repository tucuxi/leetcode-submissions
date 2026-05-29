func numberOfSpecialChars(word string) int {
    var h [64]bool
    res := 0
    for i := range word {
        b := word[i] - 0x40
        if !h[b] {
            h[b] = true
            if h[b^0x20] {
                res++
            }
        }
    }
    return res
}