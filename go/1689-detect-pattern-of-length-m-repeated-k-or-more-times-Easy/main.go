func containsPattern(arr []int, m int, k int) bool {
    outer:
    for i := 0; i <= len(arr) - m * k; i++ {
        for j := 0; j < m; j++ {
            for l := 1; l < k; l++ {
                if arr[i + j + l * m] != arr[i + j] {
                    continue outer
                }
            }
        }
        return true
    }
    return false
}