func minimumSum(n int, k int) int {
    h := make(map[int]bool)
    s := 0
    i := 1
    for j := n; j > 0; j-- {
        for h[k - i] {
            i++
        }
        h[i] = true
        s += i
        i++
    }
    return s
}