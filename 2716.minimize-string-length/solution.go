func minimizedStringLength(s string) int {
    var h [26]bool
    res := 0
    for _, ch := range s {
        if index := ch - 'a'; !h[index] {
            res++
            h[index] = true
        }
    }
    return res
}