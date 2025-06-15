func isPerfectSquare(num int) bool {
    l := 1
    r := num
    for l <= r {
        m := (l + r) / 2
        if m * m == num {
            return true
        }
        if m * m < num {
            l = m + 1
        } else {
            r = m - 1
        }
    }
    return false
}