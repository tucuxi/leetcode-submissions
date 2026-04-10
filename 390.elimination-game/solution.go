func lastRemaining(n int) int {
    return f(n, true)    
}

func f(n int, l2r bool) int {
    switch {
        case n == 1:
            return 1
        case l2r || n%2 == 1:
            return 2 * f(n/2, !l2r)
        default:
            return 2 * f(n/2, !l2r) - 1
    }
}