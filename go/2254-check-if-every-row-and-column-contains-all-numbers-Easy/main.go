func checkValid(matrix [][]int) bool {
    n := len(matrix)
    nums := make([]bool, n + 1)
    for row := 0; row < n; row++ {
        for i := 1; i <= n; i++ {
            nums[i] = false
        }
        for col := 0; col < n; col++ {
            nums[matrix[row][col]] = true
        }
        for i := 1; i <= n; i++ {
            if !nums[i] {
                return false
            }
        }
    }
    for col := 0; col < n; col++ {
        for i := 1; i <= n; i++ {
            nums[i] = false
        }
        for row := 0; row < n; row++ {
            nums[matrix[row][col]] = true
        }
        for i := 1; i <= n; i++ {
            if !nums[i] {
                return false
            }
        }
    }
    return true
}