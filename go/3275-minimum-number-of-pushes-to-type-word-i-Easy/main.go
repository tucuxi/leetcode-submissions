func minimumPushes(word string) int {
    r := 0
    k := 0
    for n := len(word); n > 0; n -= 8 {
        k++
        r += k * min(n, 8)
    }
    return r
}