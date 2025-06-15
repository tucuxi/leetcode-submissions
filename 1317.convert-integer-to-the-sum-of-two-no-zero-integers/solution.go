func getNoZeroIntegers(n int) []int {
    for a := 1; a < n; a++ {
        if nozero(a) && nozero(n - a) {
            return []int{a, n - a}
        }
    }
    panic("no solution")
}

func nozero(n int) bool {
    if n == 0 {
        return false
    }
    for n > 0 {
        if n % 10 == 0 {
            return false
        }
        n /= 10
    }
    return true
}