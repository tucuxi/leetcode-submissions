func numOfSubarrays(arr []int, k int, threshold int) int {
    if len(arr) < k {
        return 0
    }

    target := threshold*k

    w := 0
    for _, a := range arr[:k] {
        w += a
    }
    
    res := 0
    if w >= target {
        res++
    }

    for i := k; i < len(arr); i++ {
        w = w - arr[i-k] + arr[i]
        if w >= target {
            res++
        }
    }
    return res
}