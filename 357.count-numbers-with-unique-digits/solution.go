func countNumbersWithUniqueDigits(n int) int {
    switch n {
    case 0:
        return 1
    case 1:
        return 10
    default:
        c := 9
        f := 9
        for k := 2; k <= n; k++ {
            c *= f
            f--
        }
        return c + countNumbersWithUniqueDigits(n - 1)
    }
}