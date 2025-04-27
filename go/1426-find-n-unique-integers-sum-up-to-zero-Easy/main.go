func sumZero(n int) []int {
    a := make([]int, n)
    for i := 0; i < n/2; i++ {
        a[i] = i + 1
        a[n - 1 - i] = -i - 1
    }
    return a
}