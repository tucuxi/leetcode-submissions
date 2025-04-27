func maxSum(grid [][]int, limits []int, k int) int64 {
    var largest []int

    for i := range grid {
        slices.Sort(grid[i])
        largest = append(largest, grid[i][len(grid[i]) - limits[i] :]...)
    }

    slices.Sort(largest)

    var res int64
    
    for i := len(largest) - k; i < len(largest); i++ {
        res += int64(largest[i])
    }

    return res
}