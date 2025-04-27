func minOperations(n int) int {
    res := 0
    i := 0
    j := n - 1
    for i < j {
        res += j - i
        i++
        j--
    }
    return res
}
