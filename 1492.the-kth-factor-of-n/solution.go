func kthFactor(n int, k int) int {
    var a []int
    
    for i := 1; i*i <= n; i++ {
        if n % i == 0 {
            a = append(a, i)
        }
    }
    if k <= len(a) {
        return a[k-1]
    }
    if h := a[len(a) - 1]; h*h == n {
        k++
    }
    if j := 2 * len(a) - k; j >= 0 {
        return n / a[j]
    }
    return -1
}