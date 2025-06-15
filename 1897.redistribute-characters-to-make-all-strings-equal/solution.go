func makeEqual(words []string) bool {
    var h [26]int
    for _, w := range words {
        for _, r := range w {
            h[r - 'a']++
        }
    }
    for _, f := range h {
        if f % len(words) != 0 {
            return false
        }
    }
    return true
}