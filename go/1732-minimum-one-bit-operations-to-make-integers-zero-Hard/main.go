func minimumOneBitOperations(n int) int {
    res, curr := 0, 1
    for ; n > 0; n /= 2 {
        if n % 2 == 1 {
            res = curr - res
        }
        curr = 2 * curr + 1
    }
    return res
}