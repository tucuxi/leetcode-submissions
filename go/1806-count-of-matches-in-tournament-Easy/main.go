func numberOfMatches(n int) int {
    t := 0
    for n > 1 {
        t += n/2
        n = (n+1) / 2
    }
    return t
}