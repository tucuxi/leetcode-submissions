func sumOfGoodIntegers(n int, k int) int {
    res := 0

    for x := max(n-k, 1); x <= n+k; x++ {
        if n & x == 0 {
            res += x
        }
    }
    return res
}