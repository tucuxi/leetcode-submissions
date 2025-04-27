func isUgly(n int) bool {
    if n <= 0 {
        return false
    } 
    p := n
    for _, i := range []int{2, 3, 5} {
        for p % i == 0 {
            p /= i
        } 
    }
    return p == 1 
}