func trap(height []int) int {
    n := len(height)
    heighestLeft := make([]int, n)
    heighestLeft[0] = height[0]
    for i := 1; i < n; i++ {
        heighestLeft[i] = max(height[i], heighestLeft[i - 1])
    }
    heighestRight := make([]int, n)
    heighestRight[n - 1] = height[n - 1]
    for i := n - 2; i >= 0; i-- {
        heighestRight[i] = max(height[i], heighestRight[i + 1])
    }
    res := 0
    for i := range height {
        res += min(heighestLeft[i], heighestRight[i]) - height[i]
    }
    return res
}