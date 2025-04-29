func minimumLength(s string) int {
    var h [26]int
    res := 0

    for _, r := range s {
        i := r - 'a'
        if h[i] < 2 {
            h[i]++
            res++
        } else {
            h[i]--
            res--
        }
    }

    return res
}
