func canArrange(arr []int, k int) bool {
    h := make([]int, k)
    for _, a := range arr {
        h[(a % k + k) % k]++
    }
    if k % 2 == 0 && h[k / 2] % 2 != 0 {
        return false
    }
    for i := 1; i <= k / 2; i++ {
        if h[i] != h[k - i] {
            return false
        }
    }
    return true
}